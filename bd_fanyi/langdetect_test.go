package bd_fanyi

import (
	"testing"

	"github.com/rylans/getlang"
)

func TestLangDetect(t *testing.T) {
	info := getlang.FromString("Организация расположена по адресу ул")
	lang, conf := info.LanguageCode(), info.Confidence()
	if lang != "ru" || conf < 0.5 {
		t.Errorf("Got %s %f", lang, conf)
	}
	info = getlang.FromString("Wszyscy ludzie rodzą się wolni i równi w swojej godności i prawach")
	lang, conf = info.LanguageCode(), info.Confidence()
	if lang != "pl" || conf < 0.9 {
		t.Errorf("Got %s %f", lang, conf)
	}
	info = getlang.FromString("免费调用量200万字符/月")
	lang, conf = info.LanguageCode(), info.Confidence()
	if lang != "zh" || conf < 0.9 {
		t.Errorf("Got %s %f", lang, conf)
	}
	info = getlang.FromString("Access denied due to your region")
	lang, conf = info.LanguageCode(), info.Confidence()
	if lang != "en" || conf < 0.5 {
		t.Errorf("Got %s %f", lang, conf)
	}
}
