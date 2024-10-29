package types

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	common_types "github.com/EspressoSystems/espresso-sequencer-go/types/common"
	"github.com/stretchr/testify/require"
)

var (
	HEADER0_1_COMMITMENT   = common_types.Commitment{118, 29, 74, 165, 219, 239, 197, 43, 231, 156, 250, 78, 139, 108, 136, 220, 51, 160, 242, 30, 165, 182, 189, 138, 191, 93, 226, 71, 54, 208, 190, 211}
	HEADER0_2_COMMITMENT   = common_types.Commitment{87, 65, 137, 140, 189, 125, 156, 42, 229, 155, 217, 245, 205, 158, 160, 104, 226, 132, 122, 68, 140, 9, 62, 174, 71, 147, 254, 135, 177, 162, 233, 66}
	HEADER0_3_COMMITMENT   = common_types.Commitment{4, 105, 64, 105, 216, 176, 58, 92, 102, 133, 12, 93, 167, 97, 210, 238, 97, 233, 27, 232, 159, 12, 236, 125, 161, 192, 100, 76, 66, 87, 199, 78}
	HEADER0_3_2_COMMITMENT = common_types.Commitment{249, 248, 47, 186, 166, 165, 249, 214, 192, 89, 254, 52, 192, 218, 145, 40, 41, 38, 233, 185, 86, 220, 1, 125, 96, 72, 217, 46, 234, 176, 255, 25}
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

	require.Equal(t, header.Commit(), HEADER0_1_COMMITMENT)
}

func TestHeader0_2(t *testing.T) {
	header := getHeaderFromTestFile("./test-data/header0_2.json", t)

	if header.Version().Major != 0 || header.Version().Minor != 2 {
		t.Fatal("Wrong version", header.Version())
	}

	testHeaderFields(header, t)

	require.Equal(t, header.Commit(), HEADER0_2_COMMITMENT)
}

func TestHeader0_3(t *testing.T) {
	header := getHeaderFromTestFile("./test-data/header0_3.json", t)

	if header.Version().Major != 0 || header.Version().Minor != 3 {
		t.Fatal("Wrong version", header.Version())
	}

	testHeaderFields(header, t)

	require.Equal(t, header.Commit(), HEADER0_3_COMMITMENT)
}

func TestHeaderImplMarshalAndUnmarshal(t *testing.T) {
	var header HeaderInterface
	header = getHeaderFromTestFile("./test-data/header0_1.json", t)
	testHeaderImplMarshalAndUnmarshal(header, t, HEADER0_1_COMMITMENT)
	header = getHeaderFromTestFile("./test-data/header0_2.json", t)
	testHeaderImplMarshalAndUnmarshal(header, t, HEADER0_2_COMMITMENT)
	header = getHeaderFromTestFile("./test-data/header0_3.json", t)
	testHeaderImplMarshalAndUnmarshal(header, t, HEADER0_3_COMMITMENT)
	header = getHeaderFromTestFile("./test-data/header0_3_2.json", t)
	testHeaderImplMarshalAndUnmarshal(header, t, HEADER0_3_2_COMMITMENT)
}

func testHeaderImplMarshalAndUnmarshal(header HeaderInterface, t *testing.T, commitment Commitment) {
	headerImpl := HeaderImpl{Header: header}
	if headerImpl.Header.Commit() != commitment {
		t.Fatal("Incorrect commitment")
	}
	bytes, err := json.Marshal(headerImpl)
	if err != nil {
		t.Fatal("Failed to marshal header", err)
	}
	var actualHeaderImpl HeaderImpl
	err = json.Unmarshal(bytes, &actualHeaderImpl)
	if err != nil {
		t.Fatal("failed to unmarshal header", err)
	}

	if actualHeaderImpl.Header.Commit() != commitment {
		t.Fatal("Incorrect commitment after marshal")
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
