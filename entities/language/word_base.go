package language

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type WordBase struct {
	FemaleOwnNames []string `json:"female_own_names"`
	MaleOwnNames   []string `json:"male_own_names"`
	Words          []string `json:"words"`
	Name           string   `json:"name"`
	Min            int      `json:"min"`
	Max            int      `json:"max"`
	Dupl           string   `json:"dupl"`
	M              float64  `json:"m"`
}

func ExtractWords(wbs []*WordBase) map[string][]string {
	out := make(map[string][]string, len(wbs))
	for _, wb := range wbs {
		out[wb.Name] = wb.Words
	}

	return out
}

func SetWordBases() (err error) {
	SindarinWordBase, err = SindarinWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	DunmerisWordBase, err = DunmerisWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	DwemerisWordBase, err = DwemerisWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	GoblinsWordBase, err = GoblinsWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	OrcWordBase, err = OrcWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	GiantWordBase, err = GiantWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	DraconicWordBase, err = DraconicWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	ArachnidWordBase, err = ArachnidWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	SerpentsWordBase, err = SerpentsWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}

	AlbanianWordBase, err = AlbanianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	ArmenianWordBase, err = ArmenianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	RuthenianWordBase, err = RuthenianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	LithuanianWordBase, err = LithuanianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	CelticWordBase, err = CelticWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	DacianWordBase, err = DacianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	NordicWordBase, err = NordicWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	EnglishWordBase, err = EnglishWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	GermanWordBase, err = GermanWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	GreekWordBase, err = GreekWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	IllyrianWordBase, err = IllyrianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	HindiWordBase, err = HindiWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	IranianWordBase, err = IranianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	LatinWordBase, err = LatinWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	PortugueseWordBase, err = PortugueseWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	SpanishWordBase, err = SpanishWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	FrenchWordBase, err = FrenchWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	ItalianWordBase, err = ItalianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	EtruscanWordBase, err = EtruscanWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}

	ChineseWordBase, err = ChineseWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	CantoneseWordBase, err = CantoneseWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}

	JapaneseWordBase, err = JapaneseWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	KoreanWordBase, err = KoreanWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	VietnameseWordBase, err = VietnameseWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	KannadaWordBase, err = KannadaWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}

	BasqueWordBase, err = BasqueWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	GeorgianWordBase, err = GeorgianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}

	FinnicWordBase, err = FinnicWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	HungarianWordBase, err = HungarianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}

	TurkishWordBase, err = TurkishWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	MongolianWordBase, err = MongolianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	NigerianWordBase, err = NigerianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	SwahiliWordBase, err = SwahiliWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	NahuatlWordBase, err = NahuatlWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	InuitWordBase, err = InuitWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	HawaiianWordBase, err = HawaiianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	QuechuaWordBase, err = QuechuaWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	BerberWordBase, err = BerberWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	MesopotamianWordBase, err = MesopotamianWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}
	ArabicWordBase, err = ArabicWordBaseRef.LoadWordBase()
	if err != nil {
		return err
	}

	AllWordBases = []*WordBase{
		SindarinWordBase,
		DunmerisWordBase,
		DwemerisWordBase,
		GoblinsWordBase,
		OrcWordBase,
		GiantWordBase,
		DraconicWordBase,
		ArachnidWordBase,
		SerpentsWordBase,

		AlbanianWordBase,
		ArmenianWordBase,
		RuthenianWordBase,
		LithuanianWordBase,
		CelticWordBase,
		DacianWordBase,
		NordicWordBase,
		EnglishWordBase,
		GermanWordBase,
		GreekWordBase,
		IllyrianWordBase,
		HindiWordBase,
		IranianWordBase,
		LatinWordBase,
		PortugueseWordBase,
		SpanishWordBase,
		FrenchWordBase,
		ItalianWordBase,
		EtruscanWordBase,

		ChineseWordBase,
		CantoneseWordBase,

		JapaneseWordBase,
		KoreanWordBase,
		VietnameseWordBase,
		KannadaWordBase,

		BasqueWordBase,
		GeorgianWordBase,

		FinnicWordBase,
		HungarianWordBase,

		TurkishWordBase,
		MongolianWordBase,
		NigerianWordBase,
		SwahiliWordBase,
		NahuatlWordBase,
		InuitWordBase,
		HawaiianWordBase,
		QuechuaWordBase,
		BerberWordBase,
		MesopotamianWordBase,
		ArabicWordBase,
	}

	return err
}

func getWordBaseFromJSON(filename string) (*WordBase, error) {
	jsonFile, err := os.Open("entities/language/word_bases/" + filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var wb WordBase
	if err := json.Unmarshal(byteValue, &wb); err != nil {
		return nil, err
	}

	return &wb, nil
}
