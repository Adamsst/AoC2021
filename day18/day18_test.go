package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetFirstNum(t *testing.T) {
	type data struct {
		in        string
		out       int
		shouldErr bool
	}
	var tests = []data{
		{
			in:        "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			out:       4,
			shouldErr: false,
		},
		{
			in:        "[[[[[,37],4],4],[7,[[8,4],9]]],[1,1]]",
			out:       37,
			shouldErr: false,
		},
		{
			in:        "[,,,[[[[,,]]",
			out:       0,
			shouldErr: true,
		},
		{
			in:        "[,,agfd,[[gfd[f[,hfg]]",
			out:       0,
			shouldErr: true,
		},
		{
			in:        "",
			out:       0,
			shouldErr: true,
		},
		{
			in:        "[[[[[41243,3],4],4],[7,[[8,4],9]]],[1,1]]",
			out:       41243,
			shouldErr: false,
		},
		{
			in:        "[[[[[,,,65436531]]",
			out:       65436531,
			shouldErr: false,
		},
	}

	for _, test := range tests {
		res, err := getFirstNumber(test.in)
		if test.shouldErr {
			require.Error(t, err)
		}
		require.Equal(t, test.out, res)
	}
}

func Test_GetDigitsInInt(t *testing.T) {
	type data struct {
		in, out int
	}

	var tests = []data{
		{
			in:  1,
			out: 1,
		},
		{
			in:  0,
			out: 1,
		},
		{
			in:  9,
			out: 1,
		},
		{
			in:  99,
			out: 2,
		},
		{
			in:  1234,
			out: 4,
		},
		{
			in:  90909,
			out: 5,
		},
	}

	for _, test := range tests {
		res := getDigitsInInt(test.in)
		require.Equal(t, test.out, res)
	}
}

func Test_ScanToSlice(t *testing.T) {
	type data struct {
		in  string
		out []sf
	}

	var tests = []data{
		{
			in: "[4,6]",
			out: []sf{
				{
					l: sfVal{
						isSet: true,
						value: 4,
					},
					r: sfVal{
						isSet: true,
						value: 6,
					},
					d: 1,
				},
			},
		},
		{
			in: "[9,[8,7]]",
			out: []sf{
				{
					l: sfVal{
						isSet: true,
						value: 9,
					},
					r: sfVal{
						isSet: false,
						value: 0,
					},
					d: 1,
				},
				{
					l: sfVal{
						isSet: true,
						value: 8,
					},
					r: sfVal{
						isSet: true,
						value: 7,
					},
					d: 2,
				},
			},
		},
		{
			in: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			out: []sf{
				{
					l: sfVal{
						isSet: true,
						value: 4,
					},
					r: sfVal{
						isSet: true,
						value: 3,
					},
					d: 5,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: true,
						value: 4,
					},
					d: 4,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: true,
						value: 4,
					},
					d: 3,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: false,
						value: 0,
					},
					d: 2,
				},
				{
					l: sfVal{
						isSet: true,
						value: 7,
					},
					r: sfVal{
						isSet: false,
						value: 0,
					},
					d: 3,
				},
				{
					l: sfVal{
						isSet: true,
						value: 8,
					},
					r: sfVal{
						isSet: true,
						value: 4,
					},
					d: 5,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: true,
						value: 9,
					},
					d: 4,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: false,
						value: 0,
					},
					d: 1,
				},
				{
					l: sfVal{
						isSet: true,
						value: 1,
					},
					r: sfVal{
						isSet: true,
						value: 1,
					},
					d: 2,
				},
			},
		},
	}

	for _, test := range tests {
		res, err := scanToSlice(test.in)
		require.Equal(t, test.out, res)
		require.NoError(t, err)
	}
}

func Test_SliceToStr(t *testing.T) {
	type data struct {
		out string
		in  []sf
	}

	var tests = []data{
		{
			out: "[4,6]",
			in: []sf{
				{
					l: sfVal{
						isSet: true,
						value: 4,
					},
					r: sfVal{
						isSet: true,
						value: 6,
					},
					d: 1,
				},
			},
		},
		{
			out: "[9,[8,7]]",
			in: []sf{
				{
					l: sfVal{
						isSet: true,
						value: 9,
					},
					r: sfVal{
						isSet: false,
						value: 0,
					},
					d: 1,
				},
				{
					l: sfVal{
						isSet: true,
						value: 8,
					},
					r: sfVal{
						isSet: true,
						value: 7,
					},
					d: 2,
				},
			},
		},
		{
			out: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			in: []sf{
				{
					l: sfVal{
						isSet: true,
						value: 4,
					},
					r: sfVal{
						isSet: true,
						value: 3,
					},
					d: 5,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: true,
						value: 4,
					},
					d: 4,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: true,
						value: 4,
					},
					d: 3,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: false,
						value: 0,
					},
					d: 2,
				},
				{
					l: sfVal{
						isSet: true,
						value: 7,
					},
					r: sfVal{
						isSet: false,
						value: 0,
					},
					d: 3,
				},
				{
					l: sfVal{
						isSet: true,
						value: 8,
					},
					r: sfVal{
						isSet: true,
						value: 4,
					},
					d: 5,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: true,
						value: 9,
					},
					d: 4,
				},
				{
					l: sfVal{
						isSet: false,
						value: 0,
					},
					r: sfVal{
						isSet: false,
						value: 0,
					},
					d: 1,
				},
				{
					l: sfVal{
						isSet: true,
						value: 1,
					},
					r: sfVal{
						isSet: true,
						value: 1,
					},
					d: 2,
				},
			},
		},
	}

	for _, test := range tests {
		res, err := sliceToString(test.in)
		require.Equal(t, test.out, res)
		require.NoError(t, err)
	}
}

func Test_AddString(t *testing.T) {
	type data struct {
		in1 string
		in2 string
		out string
	}

	var tests = []data{
		{
			in1: "[1,2]",
			in2: "[[3,4],5]",
			out: "[[1,2],[[3,4],5]]",
		},
		{
			in1: "[[[[4,3],4],4],[7,[[8,4],9]]]",
			in2: "[1,1]",
			out: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
		},
	}

	for _, test := range tests {
		res := addString(test.in1, test.in2)
		require.Equal(t, test.out, res)
	}
}

func Test_Explode(t *testing.T) {
	type data struct {
		in  string
		out string
		exp bool
	}

	var tests = []data{
		{
			in:  "[[[[[9,8],1],2],3],4]",
			out: "[[[[0,9],2],3],4]",
			exp: true,
		},
		{
			in:  "[[[[0,9],2],3],4]",
			out: "[[[[0,9],2],3],4]",
			exp: false,
		},
		{
			in:  "[7,[6,[5,[4,[3,2]]]]]",
			out: "[7,[6,[5,[7,0]]]]",
			exp: true,
		},
		{
			in:  "[7,[6,[5,[7,0]]]]",
			out: "[7,[6,[5,[7,0]]]]",
			exp: false,
		},
		{
			in:  "[[6,[5,[4,[3,2]]]],1]",
			out: "[[6,[5,[7,0]]],3]",
			exp: true,
		},
		{
			in:  "[[6,[5,[7,0]]],3]",
			out: "[[6,[5,[7,0]]],3]",
			exp: false,
		},
		{
			in:  "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			out: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			exp: true,
		},
		{
			in:  "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			out: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
			exp: true,
		},
		{
			in:  "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
			out: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
			exp: false,
		},
		{
			in:  "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			out: "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
			exp: true,
		},
		{
			in:  "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
			out: "[[[[0,7],4],[15,[0,13]]],[1,1]]",
			exp: true,
		},
		{
			in:  "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
			out: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			exp: true,
		},
		{
			in:  "[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]",
			out: "[[[[0,[3,2]],[3,3]],[4,4]],[5,5]]",
			exp: true,
		},
		{
			in:  "[[[[0,[3,2]],[3,3]],[4,4]],[5,5]]",
			out: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			exp: true,
		},
	}

	for _, test := range tests {
		inSlice, err := scanToSlice(test.in)
		require.NoError(t, err)

		res, exploded := explode(inSlice)
		require.Equal(t, test.exp, exploded)

		res2, err := sliceToString(res)
		require.NoError(t, err)

		require.Equal(t, test.out, res2)
	}
}

func Test_StrToSliceToStr(t *testing.T) {
	type data struct {
		in string
	}

	var tests = []data{
		{
			in: "[[1,2],[[3,4],5]]",
		},
		{
			in: "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		},
		{
			in: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			in: "[[[[0,7],4],[15,[0,13]]],[1,1]]",
		},
		{
			in: "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		},
		{
			in: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}

	for _, test := range tests {
		inSlice, err := scanToSlice(test.in)
		require.NoError(t, err)

		res, err := sliceToString(inSlice)
		require.NoError(t, err)

		require.Equal(t, test.in, res)
	}
}

func Test_StrToSplitToStr(t *testing.T) {
	type data struct {
		in, out string
		isSplit bool
	}

	var tests = []data{
		{
			in:      "[[[[0,7],4],[15,[0,13]]],[1,1]]",
			out:     "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			isSplit: true,
		},
		{
			in:      "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			out:     "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
			isSplit: true,
		},
	}

	for _, test := range tests {
		inSlice, err := scanToSlice(test.in)
		require.NoError(t, err)

		splitRes, split := split(inSlice)
		require.Equal(t, split, test.isSplit)

		res, err := sliceToString(splitRes)
		require.NoError(t, err)

		require.Equal(t, test.out, res)
	}
}

func Test_PartOne(t *testing.T) {
	type data struct {
		in  []string
		out string
	}

	var tests = []data{
		{
			in: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
			},
			out: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			in: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
				"[6,6]",
			},
			out: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			in: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			},
			out: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			in: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			out: "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		},
	}

	for _, test := range tests {
		res := partOne(test.in)
		require.Equal(t, test.out, res)
	}
}

func Test_Mag(t *testing.T) {
	type data struct {
		in  string
		out int
	}

	var tests = []data{
		{
			in:  "[[1,2],[[3,4],5]]",
			out: 143,
		},
		{
			in:  "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			out: 445,
		},
		{
			in:  "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			out: 1384,
		},
		{
			in:  "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			out: 3488,
		},
		{
			in:  "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
			out: 4140,
		},
	}

	for _, test := range tests {
		inSlice, err := scanToSlice(test.in)
		require.NoError(t, err)

		res := getMagnitude(inSlice)

		require.Equal(t, test.out, res)
	}
}
