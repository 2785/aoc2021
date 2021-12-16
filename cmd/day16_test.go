package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay16(t *testing.T) {
	t0 := "38006F45291200"
	t1 := `8A004A801A8002F478`
	t2 := "620080001611562C8802118E34"
	t3 := "C0015000016115A2E0802F182340"
	t4 := "A0016C880162017C3686B18A3D4780"

	p1t0, err := solveDay16P1(t0)
	require.NoError(t, err)
	require.Equal(t, 9, p1t0)

	p1t1, err := solveDay16P1(t1)
	require.NoError(t, err)
	require.Equal(t, 16, p1t1)

	p1t2, err := solveDay16P1(t2)
	require.NoError(t, err)
	require.Equal(t, 12, p1t2)

	p1t3, err := solveDay16P1(t3)
	require.NoError(t, err)
	require.Equal(t, 23, p1t3)

	p1t4, err := solveDay16P1(t4)
	require.NoError(t, err)
	require.Equal(t, 31, p1t4)

	t5 := "C200B40A82"
	t6 := "04005AC33890"
	t7 := "880086C3E88112"
	t8 := "CE00C43D881120"
	t9 := "D8005AC2A8F0"
	t10 := "F600BC2D8F"
	t11 := "9C005AC2F8F0"
	t12 := "9C0141080250320F1802104A08"

	p2t5, err := solveDay16P2(t5)
	require.NoError(t, err)
	require.Equal(t, 3, p2t5)

	p2t6, err := solveDay16P2(t6)
	require.NoError(t, err)
	require.Equal(t, 54, p2t6)

	p2t7, err := solveDay16P2(t7)
	require.NoError(t, err)
	require.Equal(t, 7, p2t7)

	p2t8, err := solveDay16P2(t8)
	require.NoError(t, err)
	require.Equal(t, 9, p2t8)

	p2t9, err := solveDay16P2(t9)
	require.NoError(t, err)
	require.Equal(t, 1, p2t9)

	p2t10, err := solveDay16P2(t10)
	require.NoError(t, err)
	require.Equal(t, 0, p2t10)

	p2t11, err := solveDay16P2(t11)
	require.NoError(t, err)
	require.Equal(t, 0, p2t11)

	p2t12, err := solveDay16P2(t12)
	require.NoError(t, err)
	require.Equal(t, 1, p2t12)
}
