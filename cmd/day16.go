package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func init() {
	solvers[16] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay16P1,
		P2: DoDay16P2,
	}
}

type day16Packet struct {
	id         string
	version    int
	typeID     int
	subPackets []string
	content    string
	value      *int
}

func day16ParseIntoBin(input string) (string, error) {
	binStr := ""
	for _, v := range strings.Split(input, "") {
		n, err := strconv.ParseInt(v, 16, 64)
		if err != nil {
			return "", err
		}

		binStr += fmt.Sprintf("%04b", n)
	}

	return binStr, nil
}

func day16ParsePackets(input string) (rootID string, packets map[string]day16Packet, leftOver string, e error) {
	versionStr := input[0:3]
	version, err := strconv.ParseInt(versionStr, 2, 64)
	if err != nil {
		e = err
		return
	}

	pTypeStr := input[3:6]
	pType, err := strconv.ParseInt(pTypeStr, 2, 64)
	if err != nil {
		e = err
		return
	}

	packet := day16Packet{
		id:         uuid.NewString(),
		version:    int(version),
		typeID:     int(pType),
		subPackets: []string{},
	}

	packets = make(map[string]day16Packet)

	if pType == 4 {
		// we are a literal value
		end := false
		binString := ""
		lOver := input[6:]
		for i := 0; 6+i*5 < len(input); i++ {
			binStr := input[6+i*5 : 6+i*5+5]
			header := binStr[:1]

			binString += binStr[1:]
			lOver = lOver[5:]

			if header == "0" {
				end = true
				break
			}
		}

		if !end {
			e = fmt.Errorf("invalid packet")
			return
		}

		packet.content = binString
		leftOver = lOver

	} else {
		// we are an operator packet
		lenType := input[6:7]
		if lenType == "0" {
			// we're dealing with total length in bits of the subpackets
			lb := input[7 : 7+15]
			l, err := strconv.ParseInt(lb, 2, 64)
			if err != nil {
				e = err
				return
			}

			subpacketStr := input[7+15 : 7+15+l]
			for {
				h, pkts, lOver, err := day16ParsePackets(subpacketStr)
				if err != nil {
					e = err
					return
				}

				packet.subPackets = append(packet.subPackets, h)

				for k, v := range pkts {

					packets[k] = v
				}

				if strings.ReplaceAll(lOver, "0", "") == "" {
					// we're done, everything left is zeros
					break
				}

				subpacketStr = lOver
			}

			leftOver = input[7+15+l:]
		} else {
			lpb := input[7 : 7+11]
			lp, err := strconv.ParseInt(lpb, 2, 64)
			if err != nil {
				e = err
				return
			}

			remainder := input[7+11:]

			for i := 0; i < int(lp); i++ {
				h, pkts, lOver, err := day16ParsePackets(remainder)
				if err != nil {
					e = err
					return
				}

				packet.subPackets = append(packet.subPackets, h)

				for k, v := range pkts {
					packets[k] = v
				}

				remainder = lOver
			}

			leftOver = remainder
		}

	}

	rootID = packet.id
	packets[packet.id] = packet

	return rootID, packets, leftOver, nil
}

func DoDay16P1(input string) {
	sol, err := solveDay16P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 16 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay16P1(input string) (int, error) {
	bStr, err := day16ParseIntoBin(input)
	if err != nil {
		return 0, err
	}

	_, packets, _, err := day16ParsePackets(bStr)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, v := range packets {
		count += v.version
	}

	return count, nil
}

func DoDay16P2(input string) {
	sol, err := solveDay16P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 16 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay16P2(input string) (int, error) {
	bStr, err := day16ParseIntoBin(input)
	if err != nil {
		return 0, err
	}

	rootPacket, packets, _, err := day16ParsePackets(bStr)
	if err != nil {
		return 0, err
	}

	err = day16Evaluate(packets, rootPacket)
	if err != nil {
		return 0, err
	}

	return *packets[rootPacket].value, nil
}

func day16Evaluate(m map[string]day16Packet, id string) error {
	packet := m[id]
	if packet.typeID == 4 {
		num, err := strconv.ParseInt(packet.content, 2, 64)
		if err != nil {
			return err
		}

		numI := int(num)

		packet.value = &numI
		m[id] = packet
		return nil
	}

	if packet.typeID == 0 {
		sum := 0
		for _, subID := range packet.subPackets {
			if m[subID].value == nil {
				err := day16Evaluate(m, subID)
				if err != nil {
					return err
				}
			}

			sum += *m[subID].value
		}
		packet.value = &sum
		m[id] = packet
		return nil
	}

	if packet.typeID == 1 {
		prod := 1
		for _, subID := range packet.subPackets {
			if m[subID].value == nil {
				err := day16Evaluate(m, subID)
				if err != nil {
					return err
				}
			}

			prod *= *m[subID].value
		}
		packet.value = &prod
		m[id] = packet
		return nil
	}

	if packet.typeID == 2 {
		values := make([]int, 0)
		for _, subID := range packet.subPackets {
			if m[subID].value == nil {
				err := day16Evaluate(m, subID)
				if err != nil {
					return err
				}
			}

			values = append(values, *m[subID].value)
		}

		min := values[0]
		for _, v := range values {
			if v < min {
				min = v
			}
		}

		packet.value = &min
		m[id] = packet
		return nil
	}

	if packet.typeID == 3 {
		values := make([]int, 0)
		for _, subID := range packet.subPackets {
			if m[subID].value == nil {
				err := day16Evaluate(m, subID)
				if err != nil {
					return err
				}
			}

			values = append(values, *m[subID].value)
		}

		max := values[0]
		for _, v := range values {
			if v > max {
				max = v
			}
		}

		packet.value = &max
		m[id] = packet
		return nil
	}

	if packet.typeID == 5 {
		if len(packet.subPackets) != 2 {
			return errors.Errorf("invalid packet: %v", len(packet.subPackets))
		}

		if m[packet.subPackets[0]].value == nil {
			err := day16Evaluate(m, packet.subPackets[0])
			if err != nil {
				return err
			}
		}

		if m[packet.subPackets[1]].value == nil {
			err := day16Evaluate(m, packet.subPackets[1])
			if err != nil {
				return err
			}
		}

		if *m[packet.subPackets[0]].value > *m[packet.subPackets[1]].value {
			s := 1
			packet.value = &s
		} else {
			s := 0
			packet.value = &s
		}

		m[id] = packet
		return nil
	}

	if packet.typeID == 6 {
		if len(packet.subPackets) != 2 {
			return errors.Errorf("invalid packet: %v", len(packet.subPackets))
		}

		if m[packet.subPackets[0]].value == nil {
			err := day16Evaluate(m, packet.subPackets[0])
			if err != nil {
				return err
			}
		}

		if m[packet.subPackets[1]].value == nil {
			err := day16Evaluate(m, packet.subPackets[1])
			if err != nil {
				return err
			}
		}

		if *m[packet.subPackets[0]].value < *m[packet.subPackets[1]].value {
			s := 1
			packet.value = &s
		} else {
			s := 0
			packet.value = &s
		}

		m[id] = packet
		return nil
	}

	if packet.typeID == 7 {
		if len(packet.subPackets) != 2 {
			return errors.New("invalid packet")
		}

		if m[packet.subPackets[0]].value == nil {
			err := day16Evaluate(m, packet.subPackets[0])
			if err != nil {
				return err
			}
		}

		if m[packet.subPackets[1]].value == nil {
			err := day16Evaluate(m, packet.subPackets[1])
			if err != nil {
				return err
			}
		}

		if *m[packet.subPackets[0]].value == *m[packet.subPackets[1]].value {
			s := 1
			packet.value = &s
		} else {
			s := 0
			packet.value = &s
		}

		m[id] = packet
		return nil
	}

	return errors.New("invalid packet")
}
