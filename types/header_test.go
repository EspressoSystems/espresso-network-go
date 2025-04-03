package types

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	common_types "github.com/EspressoSystems/espresso-network-go/types/common"
	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	s := `{"Version":{"major":0,"minor":3}}`
	data := []byte(s)
	var v common_types.Version
	err := json.Unmarshal(data, &v)
	if err != nil {
		t.Fatal("Failed to marshal JSON", err)
	}

	if !(v.Major == 0 && v.Minor == 3) {
		t.Fatal("Get the wrong version", v)
	}

	bytes, err := json.Marshal(v)
	if err != nil {
		t.Fatal("Failed to marshal version", err)
	}
	var a common_types.Version
	if err = json.Unmarshal(bytes, &a); err != nil {
		t.Fatal("Failed to unmarshal version", err)
	}
}

func TestHeader0_1(t *testing.T) {
	header := getHeaderFromTestFile("./test-data/header0_1.json", t)

	if header.Version().Major != 0 || header.Version().Minor != 1 {
		t.Fatal("Wrong version", header.Version())
	}

	testHeaderFields(header, t)

	require.Equal(t, header.Commit(), common_types.Commitment{118, 29, 74, 165, 219, 239, 197, 43, 231, 156, 250, 78, 139, 108, 136, 220, 51, 160, 242, 30, 165, 182, 189, 138, 191, 93, 226, 71, 54, 208, 190, 211})
}

func TestHeader0_2(t *testing.T) {
	header := getHeaderFromTestFile("./test-data/header0_2.json", t)

	if header.Version().Major != 0 || header.Version().Minor != 2 {
		t.Fatal("Wrong version", header.Version())
	}

	testHeaderFields(header, t)

	require.Equal(t, header.Commit(), common_types.Commitment{87, 65, 137, 140, 189, 125, 156, 42, 229, 155, 217, 245, 205, 158, 160, 104, 226, 132, 122, 68, 140, 9, 62, 174, 71, 147, 254, 135, 177, 162, 233, 66})
}

func TestHeader0_3(t *testing.T) {
	header := getHeaderFromTestFile("./test-data/header0_3.json", t)

	if header.Version().Major != 0 || header.Version().Minor != 3 {
		t.Fatal("Wrong version", header.Version())
	}

	testHeaderFields(header, t)

	require.Equal(t, header.Commit(), common_types.Commitment{32, 245, 22, 153, 110, 177, 6, 208, 214, 120, 161, 164, 229, 62, 89, 80, 6, 26, 216, 232, 7, 137, 13, 22, 211, 123, 166, 253, 29, 22, 225, 61})
}

func TestHeaderImplMarshalAndUnmarshal(t *testing.T) {
	var header HeaderInterface
	header = getHeaderFromTestFile("./test-data/header0_1.json", t)
	testHeaderImplMarshalAndUnmarshal(header, t)
	header = getHeaderFromTestFile("./test-data/header0_2.json", t)
	testHeaderImplMarshalAndUnmarshal(header, t)
	header = getHeaderFromTestFile("./test-data/header0_3.json", t)
	testHeaderImplMarshalAndUnmarshal(header, t)
}

func testHeaderImplMarshalAndUnmarshal(header HeaderInterface, t *testing.T) {
	headerImpl := HeaderImpl{Header: header}
	bytes, err := json.Marshal(headerImpl)
	if err != nil {
		t.Fatal("Failed to marshal header", err)
	}
	var actualHeaderImpl HeaderImpl
	err = json.Unmarshal(bytes, &actualHeaderImpl)
	if err != nil {
		t.Fatal("failed to unmarshal header", err)
	}
}

func testHeaderFields(header HeaderInterface, t *testing.T) {
	if header.GetBlockHeight() != 42 {
		t.Fatal("Wrong block height", header.GetBlockHeight())
	}

	if header.GetBuilderCommitment().String() != "BUILDER_COMMITMENT~jlEvJoHPETCSwXF6UKcD22zOjfoHGuyVFTVkP_BNc-no" {
		t.Fatal("Wrong builder commitment", header.GetBuilderCommitment().String())
	}

}

func getHeaderFromTestFile(path string, t *testing.T) HeaderInterface {
	file, err := os.Open(path)
	if err != nil {
		t.Fatal("failed to open file:", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatal("Error reading file:", err)
	}

	var headerImpl HeaderImpl
	err = json.Unmarshal(data, &headerImpl)
	if err != nil {
		t.Fatal("Error unmarshaling:", err)
	}

	return headerImpl.Header
}

func TestUnmarshalSignature(t *testing.T) {
	// `r` ans `s` are hex string of odd length.
	// It should be unmarshalled successfully
	data := `
{
    "r": "0xa1c",
    "s": "0x202",
    "v": 27
}
	`
	var signature Signature
	err := json.Unmarshal([]byte(data), &signature)
	if err != nil {
		t.Fatal("Error unmarshaling:", err)
	}
	expectedR := int64(2588)
	expectedS := int64(514)
	expectedV := uint64(27)

	if expectedR != signature.R.Int64() {
		t.Fatal("getting a wrong r in unmarshal signature")
	}

	if expectedS != signature.S.Int64() {
		t.Fatal("getting a wrong r in unmarshal signature")
	}

	if expectedV != signature.V {
		t.Fatal("getting a wrong r in unmarshal signature")
	}

}
