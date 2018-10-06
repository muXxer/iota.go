package utils

import (
	"github.com/iotaledger/iota.go/trinary"
	"testing"
)

type testcase struct {
	value  trinary.Trytes
	length int
	expect bool
}

var testcases = []testcase{
	{value: "ABCDEF", length: 6, expect: true},
	{value: "ABCDEF", length: 5, expect: false},
	{value: "abcdef", length: 6, expect: false},
	{value: "abcdef", length: 5, expect: false},
}

var testcases2 = []testcase{
	{value: "ABCDEF", length: 10, expect: true},
	{value: "ABCDEF", length: 6, expect: true},
	{value: "abcdef", length: 6, expect: false},
	{value: "abcdef", length: 5, expect: false},
}

func TestIsTrytesOfExactLength(t *testing.T) {
	for i, test := range testcases {
		match := IsTrytesOfExactLength(test.value, test.length)
		if match != test.expect {
			t.Fatalf("#%d expected match for value %s to be %v but was %v\n", i, test.value, test.expect, match)
		}
	}
}

func TestIsTrytesOfMaxLength(t *testing.T) {
	for i, test := range testcases2 {
		match := IsTrytesOfMaxLength(test.value, test.length)
		if match != test.expect {
			t.Fatalf("#%d expected match for value %s to be %v but was %v\n", i, test.value, test.expect, match)
		}
	}
}

func TestIsEmptyTrytes(t *testing.T) {
	reallyEmpty := trinary.Trytes("99999999")
	if !IsEmptyTrytes(reallyEmpty) {
		t.Fatalf("expected value '%s' to act as empty\n", reallyEmpty)
	}

	notReallyEmptyExamples := []trinary.Trytes{
		("999 999"), (""), (" 999999"),
	}
	for _, notReallyEmpty := range notReallyEmptyExamples {
		if IsEmptyTrytes(notReallyEmpty) {
			t.Fatalf("expected value '%s' to not act as empty\n", notReallyEmpty)
		}
	}
}

var validAddresses = []trinary.Trytes{
	("TXBGJB9NORCEHAAWVCQRC9GQSLQCWUIKDOBYTDKVYY9GUQHPJQMKHGNWRWIFLEBPJNAAIOMUFRFLDQUECB9UMGFVBD"),
	("TXBGJB9NORCEHAAWVCQRC9GQSLQCWUIKDOBYTDKVYY9GUQHPJQMKHGNWRWIFLEBPJNAAIOMUFRFLDQUEC"),
	("INLF9FQCR9XWEKPKLUPZZJMRRDUGSOCKWVYFIDMHEVKKZIKJHGWPNMCFQ9KOONDHBOOHYWSPCZMMMKKPYFETUFYMIX"),
	("INLF9FQCR9XWEKPKLUPZZJMRRDUGSOCKWVYFIDMHEVKKZIKJHGWPNMCFQ9KOONDHBOOHYWSPCZMMMKKPY"),
}

func TestIsHash(t *testing.T) {
	for _, validAddress := range validAddresses {
		if !IsHash(validAddress) {
			t.Fatalf("expected %s to be a hash\n", validAddress)
		}
	}
}

var txHashesMWM14 = []string{
	"MPDUQUZCQANRQPYTVROTXUGWSDESYZPMIMOGCGPX9IAWMXIPTHBIDNRJFYFQKCZWUVNYCQQGGX9OA9999",
	"IKMAFVAK9IMRAM9OYIJTMYTMCBGRHTJZ99RMJMVRELMULVSWRMXRJDSEG9PZJUARY9ESH9ARVXLA99999",
	"KVCLRTNBXSHSTVXPRYLWIWFH9WCDXCA9UUL9KKXJTNII9UFBX9ZQTVK9FVHQEWGZZEMQRLEDULTBA9999",
}

func TestIsTxHashWithMWM(t *testing.T) {
	for _, txHash := range txHashesMWM14 {
		if !IsTxHashWithMWM(trinary.Trytes(txHash), 14) {
			t.Fatalf("expected a MWM of 14 to be correct for transaction hash %s\n", txHash)
		}
	}
}

var txTrytesMWM14 = []trinary.Trytes{
	("9IWJLCSRVTZTDASICSPUVKGJZZEHBMATMWBV9YMBWYSLKKSSBNSSQDNGUP9ZHIFSHUPESIIGRQTSXAW9I999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999EVXMEYTKWRWVEIMNOFLHXXONGIPFCFSICLRWCNWCRKJQPFRFOWOFSGWFGVMMFBTQHTOVEMBHMSMDXBYPDXQVPAVG99999999999999999999999999999999999999999999999CG9SCZD99C99999999C99999999BULLKAERTFLINNSOAVJKLLTWQDKQULLEJMEHCFLDLFWKKRLRPI9HVG9FYWGRPAJLLQXTC9CYFRWEIBLNDLZWOXOKLZOAWURPJMHPZJTDKNCLXQJYCUQJXEPKLXOKLTYVVSOZZDUROBGQXDLYGSRYGIQQIDCPWZ9999QEXXXIRVOQMLQLOMCJWSZKWTKOSZUZAHSZYQV9BZONGKCGIJYPLYKTXKC9SPXQESYSAIOJCLUQCG99999999999999999999999999999999EDUBDRYLE999999999MMMMMMMMMFUDTQBC9DI9YDTUBFBVMKPMJVEH"),
	("XIPBSOCWQEKXZGPVPBGJZZGCNKPEQTKUYFQY9RXJKFFXMSKLG9TGANVU9WXJDMUULBSQADSNLBFP9RCCCPQEWKZSUQDWKCHB9KYKVGXCFEETOAUZFOTMSJIMEMPDYHBNVOWNSEPQLSQMFKHSDFSVWNKPKLMJXXHAFCSESQGSRQLJEMRYQ9FSJVTOBWBOIA9RMEOGXTKCYWHEBNQXFRPOKNJOFQLUIQQSYAPODSIKCRPRXFDWESZTFQTARFADZAHYPILD99IISHEDUDIASIXXSY9NHHDJDCYCBOINAD9FZDWXODSTTPMVYUFBVFXZGYVV9ZMWBWZPRR9LTRODPBHCAJ9RZRKMGEFKTRPDWZCZP9ZRETASHWFUOIHUUTTELEHLRJY9LNXKZKMZNXDBEWOGCRLZ9SAC9ZWNRCXLOPHPONLMA9XZRWGBMIOHDDWDUNWAMYZB9WTUUBSNSRGA9IIMHXXCJVMEJYNZBDYONAUWJHCPGFIEMCZILOFQNGUGD9WRVSDCUZHFQX9ADSOHYBNSZVOQTIIYUEJIZPEWAPGWBIARPHCQKMXGMVYLEODSMYD9BAOTPOY9NYB9ORNZTTAAAPWFSPZCDMYDPJALHHWWYFAKGVYIMPNIPZLCVM9AQGFNQIMMZXE9ESKUOBGTZCOFPTIGXNHVHLKPDZAQLLANFUOZRZNMPIARNEGMRATOLTSVJKPJGDYAFXACWDYEBSDNNZQ9AOM9DTVQYQDTNHNIHDVPHLPWAIQHBOKLQHJYZQCXJULXQKURALFSLAPYLYQQDEIGUUMWPTITTEKABLGPA9OBAFS9SZPFAULOHJP9BRWOWACXTZGNFXXBWGEVMMMDGBNZTEJACBASE9SPHEHOOWAYCXRVL9ZKI9RP9PDENITLEAESPLI9KIOQGGXJFV99KUVAZHSSHKCQOASHZTDINYQVIICPXRPUAQIJFIDMPKVNDRVAR9HWSSUDWGWCWMR9JIGPCRYGUBEJVVUBQFZOQZT99JROTTTKHDDGSYOUEJGPZEACTYTUZVZYNTIMESZDISPAIYHQWXYJEELE9ZYMHTNKHUHXUPBRGFKCFQMGPOWUMTANYQJYFQVDSXDMJYTOJBVHYBUNJKZQMNWEKEICNWDLIDLCXYFHEQTGOADY9ULLISUYOJRPCDCSNTLQ9GPPHKNBOHKCUVWPOPISHQHSAKUAFFWCLMWFDZUGXYAIUPWDWTUDRVQCGNJKJGLHG9ZJJHLPAFBAAFECLXMOWTHYEFXCACFCJTFVQIDVBHUQBFGL9PT9EVSZIALUJVBBP9QGKSPAXCTEPOYBQELVTSWDOSPGWKJTUPOSJDLAEKMIUQROIDGETNRLJDKWHAQMYUYUTVWIACEZDNPCXBFRDODASAPLTGUSEGJCJUNLQITGADYFJZEJLBQFV9RBMULSCKOL9VORSVFCPKYKPHNRDHIZCYBIASDCTDIDQOEAEKYYQVFOLEBZTNUI99KYNSBEGDVYYNZSZQJFDACGYFOIPTDARXUDJEJQUIJBOBEFAMONLKPFDDDQCONJLNQJWGUPBXTLJCTUYQJPYEOXCKYHITOBOJEFSFWFSLPYCIUWHXFCTVPMSPEUODMFFEOSXCQAJRFDDAGOM9FQTTTJPOE99JUCBETHNBJQFSPNFTVTEJIYAZRJWOTSAIBLPCHSOINGSSNYCMCUXPCEBPRCRERSZTOVZLCVNTJGDEEMAOENJYNE9Y99IVKSVVUS9IOEYTWFGIJA9YNPKCCMQETOKKLKAEDFCNMPNHBOXDCDDWUJLCFFEIWXNXDJEJSLUDVFBMFOKWD9KPIBUYAQXVCYDJMRYMXHKBSUXPWIYOKDILLNCVFHKAXY9KQFFMCXXKILRPHPG9NXMEFOCBPYYGFKIXQUPSUQV9CXJVTKVNULILQOTNVXMZXIBRHBEGZOKX9QJDEYFMHPXOYAOUVZGYYJPWPFDFJBCFE9MOTMXPZNTEMRKIACIIBQTAOYUMYMUQXJIDXIQKGSRZVAJKZFVVTORPELSSPXXIEVQW9UAXIHVLVXQSFNZPRMNWFJBKQMRZHX9TDXMPFFMZEUTYAJGWF9ZBKMCIRAHUBVMOLLHQEUQNXJYXGGXBDVTPCLIKSUXEXCTFJSHUMKBBZGKYQYQSRMYOXXOMABYVVCXXDBQTIVTBUDDTLEXZDIBSLZL9DLJVD9HECYTN9HFZDJUEEWGIZDOZTLTAIPAGIEWTBOTIYOFCNRHANHABHTL99ARAIN9MT9PZVQQVLCHFFQHWADWMLFY9999999999999999999999TRAVELING999IOTA99999999999LHCSCZD99A99999999B99999999WCDDSBZGCWOJJGKOOGCTHIEBTKD9VLEKAXLPALBKWNCLXIXHSEOFJJTYHIESVLHIEFNHDSCCQTQJKGEH9JYKG9WTSMZJJTMAST9WOXMOTBZTZYWAXOQHEPIJPMZEZXWRHYUIXWLA9XMZMYTSGKJXDN9AIKQH9A9999OLGHITHAWQDTYZHDDZGHEBUULXDWC9ZMQRLXSWHWZYNQV99LPJDSUZZX9GSUOLXAQFFLHIZ9MZUAZ9999TRAVELING999IOTA99999999999CHUGHRYLE999999999MMMMMMMMMPIDIVER9GD999999UTVVKNVMMMM"),
	("999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999FZMHAINEO9LZLDZXEBU9VULUGLY9XS9YTHONDNMLQWTFLMTMWWRZFNHQ9DMFJZHLNSJDIYARUKPYKIFHXNWB999999999999999999999999YRAVELING999IOTA99999999999NKCSCZD99999999999B99999999XSMUNBIAO9IGBLAHDJDCTOTNLEHJEJPIXPYRAKU9DJJCRPQIELBNVYHHBFOBZQZREFGREITDILOSGSIOW9IRE9NYSKSIBBXAPHPWRBQYLDEKYDY9XVUFZZDYMVTBJQGCC9OWLJYNTGNABSBGHNQUDBGMAZCTEZ9999WXXGIQLRPZCCECTHPIHCPMZQUWWGDFQHCJBQPFXQYBGQBHIDCYDQAIYRKBJHGXGZWXAMTDONQMTAZ9999TRAVELING999IOTA99999999999UPRJHRYLE999999999MMMMMMMMMPIDIVER9KY999999OKKGMVKMMMM"),
}

func TestIsTransactionTrytesWithMWM(t *testing.T) {
	for _, txTrytes := range txTrytesMWM14 {
		ok, err := IsTransactionTrytesWithMWM(txTrytes, 14)
		if err != nil {
			t.Fatalf("didn't expect error %v from test values\n", err)
		}
		if !ok {
			t.Fatalf("expected a MWM of 14 to be correct for transaction trytes %s\n", txTrytes)
		}
	}
}

func TestIsAttachedTrytes(t *testing.T) {
	for _, txTrytes := range txTrytesMWM14 {
		if !IsAttachedTrytes(txTrytes) {
			t.Fatalf("expected transaction to count as attached for transaction trytes %s\n", txTrytes)
		}
	}
}
