package question

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestTheLengthOfTheLastWordInTheString(t *testing.T) {
	params := "hello world\n"
	value := TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "   hello world\n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "hello   world\n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "hello world   \n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "   hello   world   \n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 5 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 5, value)
	}

	params = "hello-world\n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 11 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 11, value)
	}

	params = "\n"
	value = TheLengthOfTheLastWordInTheString(params)
	if value != 0 {
		t.Fatalf("params: [%s], correct: %d, your: %d", params, 0, value)
	}
}

func TestCountTheNumberOfCharacters(t *testing.T) {
	param1 := "ABCDEF"
	param2 := "A"
	your := CountTheNumberOfCharacters(param1, param2)
	correct := 1
	if your != correct {
		t.Fatalf("param1: %s, param2: %s, correct: %d, your: %d", param1, param2, correct, your)
	}

	param1 = "ABCDEFAABDD"
	param2 = "A"
	your = CountTheNumberOfCharacters(param1, param2)
	correct = 3
	if your != correct {
		t.Fatalf("param1: %s, param2: %s, correct: %d, your: %d", param1, param2, correct, your)
	}

}

func TestObviouslyRandomNumbers(t *testing.T) {
	param := []int{3, 3, 2, 4}
	result := ObviouslyRandomNumbers(param)
	t.Logf("%v", result)
}

func TestBase(t *testing.T) {
	res, err := strconv.ParseInt("ABC3", 16, 64)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%d", res)
}

func TestSortStrings(t *testing.T) {
	stringArray := []string{"cap", "to", "cat", "card", "two"}
	sort.Strings(stringArray)
	t.Logf("%v", stringArray)
}

func TestWeighingWeight(t *testing.T) {
	weight := []int{1, 2}
	count := []int{2, 1}
	correct := 5
	your := WeighingWeight(weight, count)
	if your != correct {
		t.Fatalf("weight: %v, count: %v, correct: %d, your: %d", weight, count, correct, your)
	}
}

func TestMazeProblem(t *testing.T) {
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
	}
	MazeProblem(maze, 0, 0)
}

func TestTheBeautyOfTheName(t *testing.T) {
	name := "zhangsan"
	correct := 192
	your := TheBeautyOfTheName(name)
	if your != correct {
		t.Fatalf("name: %s, correct: %d, your: %d", name, correct, your)
	}

	name = "lisi"
	correct = 101
	your = TheBeautyOfTheName(name)
	if your != correct {
		t.Fatalf("name: %s, correct: %d, your: %d", name, correct, your)
	}

	name = "lcfynwnkcttqmmuicxjofgqnsdyymktpvmgevfspnnjlcphdgmzizzubhfmqnnsmbfvvspmhhukxlfmsrolfmobyfdbhdgzvktjvptjhxchntcdmumxyqalhnelsspnwayszauakstpgkbjjiospjszahxyixvkmtukqgrsrufirtslysjbhjzreverdfggnlniojtndpyoflqgsnsihjqrlwxbqsziurivumalolhsisdvkphpsctcntiuvzcrzuxeenbizwhqhlzqindrtlraqslytozoeloavgidtofshyzchgtqxsmcpoqvyvzaunapgfkqusldlgugyqfckrjzdrhqabdwrihbmfszcwgtjpldwmclivucigrpsatnmnjbruivgderdsmvqxmackqapoogzeuecwgyrvqmdyuscugwhvtrewnvxqjstvhldwgauteggfqmhnbubladbltfppxefjhcckfepybsetmrfokovuqolqrfjlpwpzizbacraprigcwlsgavbroggxyqlgzcejeftosizozwxtppjtzakqoyfxsgvlmcnkinvgntmuxdoujmtjnqoaxqzhmpvjuborilapgmcgkbunshbbqpjnoescmwoyeyhdaishaqlgvceytmdbkdwpjsonswkwkgecwcdmlsfgipxutkgmktboqbluffkbqfkhqbqadhsaeilfyvbiumkekivesdrpradoeoodgvmwjtqthswngopyvsdhmkybpkxcwsfwjsxcqijvyzhzdglswjbxyuhxrkciurkpsriljgzmemhzzkjpdtkpdndrcrqlxalnxsmwxuijxnkphmvgqgodfcfapqercwmfzeiublwjbhvakumwxiowomslmskthndblpcysufdmxrwktfjseiubkydmnweszodezmwokwalfmpgemnsxrpimomzectpqtketkojtxsavuhjmihwijuepjpytatybkzvnmkyakvkityrjbbjipnlhsdfvzhjhavsupvrdkdchrjvudjhtcedgzvtphabitivwzroakwxprfjjszidyzhnryfxjiuolwxyjkkmrajdcutvxofnpuxefsigsrdieaqdnyhebpsuwmgpfkhjxkcczbdticpdxxezmtpygklgruenftcdjnodzeayajhtusbmafrfethglzqzgwtwevaqdsxydcfvnfweqtrdkwweeyvtruxnihulyvqjgnurmbnseepobowaeywhktiomctuugjvykmmjjxdbydxwjjedwttxbzttxqhmkeuiqiwgahsegrxlmjagaekdqfpjxjjatwsfgcpqgjbpuiilvmgedevaasjvecqxhrqvwzfsrhrekdwcpgbqbfpdkueaopwmkxojkktsddbvnkzhsmppmzjbrrohqtvfsvatlybqgufgajcronrobiwkgrrzavkhhuqpdfeyorloqfragubpqtyodztnbzikqzqjwcrptsrmqizymkaxmmedsfzxxhbchhjkysvhfjjhensslbfezdbezwdlgrzjdqlmnnqlpafhqllvhjoecuklvndxmxmrwgezflzfegirwhxlvzafjvjzwpicpemqzjcxtydaxsjszykzdwgiwixruhcjpfaielqfcmtfrklzridtcdgqegtvohubqqjmjshdvdbcixbgzwxlrclnjahqwqciqtgkecbpaehwyjyphcyekikgu"
	correct = 24900
	your = TheBeautyOfTheName(name)
	if your != correct {
		t.Fatalf("name: %s, correct: %d, your: %d", name, correct, your)
	}

	name = "wqieoyctanxxubqajvuuylggeozeihdvuzpkmimizpmhxuvrivhjxktvzqkzhrazgzddbtyadzsmeztgssbvvbcjntoovwgppgefinjvtozmmgfgytaoqeixqaxnsrkfdskztoaqapqucjfuseighwecxdbhfeggeztblcufyjqtelxlmzhvvhlnwmtftaltfzbwmssjthaturgkzjqucapimlrpynxaziiyluqcghblgpfuifglztlxpypcfrockkpwfhbuegprvqdyeuefvjyksvrclgdjoxsojedwbkpmkidnsvtytlmavutbetujpbvsxyqffucgmetsunbarbmtfqzakblimdmqhjfstertnlgkqyxqfdshjitptvlkryuppjbanrcfwisquogjowjdakgtswoubvmbrychjlbwoeepyadhtrkadungtzgijbpbkjqcdartpuqproecbzziqvtahqisbfuoyhdtefbsopmsinmcdnswrugcloayjudayjldysxnqpgqrwvimiiojsbhlfyuerthujqiuvgejseklfjbbfhzyzexulwhzmpdyplkophfhscakdccvrvcrwkwhtrhapyigqyjcamwmgomoeoblpceupzcinbiclzivvlbdtvunxyjqzmvcmfagfnqngjbuwfgbqzdmwowdnttfgyseouxtinbgfzjsjrzpgxdhhnffuervkkylahzoibsanreovnxwvuzcarnetoxbwawxvjbxroytcanfzdnwoblmtdbpugaglipzfgvbccewnyluhstgiejdkuifksaoxtylbscsbqtoriaizxciwxtzznragbmntgmptbitsldqybgckkopdituvhlwgfmbjqfmjmtbwrkruukniuwkauinlejeeoqlxzjfvxiwqrmzkbzocjhgxnfbkhxjcirsyjvwclxqhrskvwpahhgjwnjlkpzozdxupzgbymxsfuspblaskerisrcjoqwjoscrcuervpdrqggnhxctjvmvcmhlnhayjsycxsbnnxqxarsqeefnehcfkxghiggspjwnseprulhqggxyibmdrovkqftummwierdfliturxxwveulxfudbrvfuiqtvnxuiceplcynmnvtfwosarbdrdnfpsysnqnhgzfamxtiozdsgntiuomxmmltqxjsxrgdqjmgqzofnuuxxjffvcfmzvdcwjbatxksbwbcgoxdyxsxeqguayyhwiqcqgwzmwbmzazxwedbohsbhryvcgdeqfmcbrpyzzcjmapdranonkknaozwcgqzcihcqbmifnpthkcxnflyfqnvwitcrfustxebobmayorvwudxidybirhghoiqvlrsaluvkcygdqdlqbdappljojetwtqsvwdxgofquflbklxznglaxhlbzshupkohnekngowwsdtlwtmwvptoznceuqvmpjlyxlamidjfsmhhyrwqyvvkowrwduocemgxwtemqkjsbuovogqqxrobbryerrxgfppbguyismaxhaurkqfasyokvcfysjmidpwmagvhrnthslzmzpqqfbcybexpebgyfwlecbdmsgkoweyguknibgkmxdflphtnafyvqprvpcxzcytjsbxurowrmzejunrhdubgprmpcbszvxrnnlpdhddplpnnqklysdqkgtpdaqxbrukarmpgnoadlywbsrvsdpjojirctsmklmgqjtaoaxuwpxibnvlqmcqeohmdsuinbhkntykqlxltdvduxeryrdrzexqqliklgulvzheacmspuugxpomagimaqnrsoqwomgvmbgvavdxzbzensgyvihupo"
	correct = 26260
	your = TheBeautyOfTheName(name)
	if your != correct {
		t.Fatalf("name: %s, correct: %d, your: %d", name, correct, your)
	}
}

func TestTruncateAStringByByte(t *testing.T) {
	param := "我ABC汉DEF"
	count := 6
	correct := "我ABC"
	your := TruncateAStringByByte(param, count)
	if your != correct {
		t.Fatalf("param: %s, count: %d, correct: %s, your: %s", param, count, correct, your)
	}
}

func TestPrint(t *testing.T) {
	fmt.Print("fafds")
}

func TestLegalIP(t *testing.T) {
	param := "10.244.2.3"
	correct := true
	your := LegalIP(param)
	if your != correct {
		t.Fatalf("param: %s, correct: %v, your: %v", param, correct, your)
	}
}
