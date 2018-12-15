package main

import (
	"fmt"
	"strings"
)

func main() {
	num2Total := 0
	num3Total := 0
	for _, id := range strings.Split(input, "\n") {
		idMap := make(map[rune]int)
		for _, c := range id {
			idMap[c]++
		}
		num2 := 0
		num3 := 0
		for _, v := range idMap {
			if v == 2 {
				num2 = 1
			}
			if v == 3 {
				num3 = 1
			}
		}
		num2Total += num2
		num3Total += num3
	}
	fmt.Println(num2Total * num3Total)
}

const input = `krdmtuqjmwfoevnadixyclzspv
yrdmtuqjiwfoevnabfxyclzsph
kqjvtuqjgwfoevnabixyclzsph
krdmtuqjgwjoevnaolxyclzsph
krdmtnqjgwfoevnabiiyxlzsph
lrymtuqjgwhoevnabixyclzsph
krdmguqjgwfoevnabixkclzsah
krdmtuqjgwfoevnibinyclzdph
krdmtucjgwfoevnabhxyclzspv
krdmtuqjgwfoevtabixyulzsuh
krdmtuqqgwfoevnabixdblzsph
krdmtuqjawfsevnabiyyclzsph
krdmtuqjgwfoevnabzxccldsph
krdmtcqegwfhevnabixyclzsph
krdmtuqjgwforvnaxixycgzsph
krdmtuqjgwfoqvnaxixyclzskh
krdmtutjgwfoevyajixyclzsph
krdmtuqmgwfoevnabixycxzspc
krdptuqjgwhoevkabixyclzsph
krdttuqjhwfoevnabixyclzspa
krdmtuqjgwfoevnabibyhnzsph
krdmtuqjywfoevntbidyclzsph
krdmtojdgwfoevnabixyclzsph
krdmtuqjgpfuevnauixyclzsph
krdmtoqjgwfrevjabixyclzsph
krdmtuqjgwfoyvndbixyclzyph
krdmtxqjgwfomvnayixyclzsph
crdmtuqjgwfoevnabixyoxzsph
krdmtpqjgwfdevnabixycqzsph
krdmtuqjgwfoevuabfxsclzsph
krdmtuqjgwfoevnybixycdzskh
krdmtusjgwfoevnabixxclzdph
krdmtuqjgwfoevnaboxyglzjph
zrdmtuqjgrfoevnalixyclzsph
krdmtuqjclfoevnabixyclzsih
kqdmtlqjgwfoevnabtxyclzsph
krdmtuqggwpoevnabixyclzlph
krdmtuqjgwfobwnrbixyclzsph
krdmtuqjgwfoevwabkxycnzsph
kldmtuqjgwfogvyabixyclzsph
krdmtuqvgwfoevnabixtcrzsph
krdmtuqjgwroevnabixyrlzspw
krdmtuqjgjfoevnabixyelzrph
krdmtuqjgffoevnaaixyclzspa
krdmtuqjgwfoevxabifywlzsph
krdmtuqjgwfoevlabixycrzsrh
krdmtuqjgwfpevnabixocqzsph
krdmtuqjgwfoevdabixycnhsph
krdmtmqjgwfoevnajixyclvsph
krdmtuqjjvfoevnabgxyclzsph
krzmtuqjgwfoevnabioyckzsph
kodmtwqjgwfoevnabieyclzsph
ehdmthqjgwfoevnabixyclzsph
krdmtuqjxwioevnabixyclbsph
grdmkutjgwfoevnabixyclzsph
krdutuqjgwfoebnabixaclzsph
krdmtuqjgwfoebnabixyclcjph
krdmteqjgwfoevnlbixycizsph
krdmtegjgwhoevnabixyclzsph
krdmtuqjgwfdrvnabixbclzsph
krdmtuqjgyfoevidbixyclzsph
krdmtubjawfoevnabixyclzuph
krdmtuqjgwfoavjabixyclzssh
krdmtuqjgwfoeonabixyclzsvo
vrdmtuqjgwffevnabixpclzsph
krdmtuqonwfoevnabixycfzsph
krdmtumjgwfpevnubixyclzsph
krdmtutjgwfoevnaciyyclzsph
krdrtuqjgwfoevnwbixyglzsph
krdmtuqjgwfoevbabixyclesdh
krdmtuqcgwfoevnabixyqdzsph
krdmtuqjgwfogvnabrxycezsph
krdmujqkgwfoevnabixyclzsph
krdmtuqjgtooevnabixyclzzph
jrdntuqjgwfoevnabixyclrsph
krdmtuqjgzfoevkebixyclzsph
krdmtuqjgwfosvnaeixyclztph
krdmtuqjgwfoevzabixydlzaph
krdmtuqzgwfoavnabiqyclzsph
krdmtuqvgwfoevnabixycwzspv
krdmvuqjgwteevnabixyclzsph
krdmtujjgwfoevgybixyclzsph
kydmtuqjgwfoeunacixyclzsph
krdmtuqjgifoqvnabicyclzsph
krnltiqjgwfoevnabixyclzsph
krdmtuqjgwfoevnabhxyclzsgi
kfdmtuqjnwfowvnabixyclzsph
kmdmtuljgwfoevnabixycvzsph
krdmtxqjgwaoevvabixyclzsph
kramduqjgwfoevnabixyclzwph
krdutuqjgwfoennabixyclziph
krdmvuqfgwfoevnacixyclzsph
krdmtuqogwfoevnabmvyclzsph
krdmfuqjgwfoyvnabixyclzseh
krdmtuqjgweoelnabixyclzspd
krdmtumjgwfoevnabixyclzypo
krdmtuqjgkfoevhabixyclzsqh
kjdmtuqjgwfoevgabixyclzsah
krdmtuqjgwfoevnlbixyclzsbw
mrdmtxqjgwfoevnabgxyclzsph
krdmtuqpgwfoevnhbixycltsph
krdmtuqjgwfmqvnabixyclzslh
krqmtuqogwfoevnaqixyclzsph
krdmtusjggfoevnabicyclzsph
krcmtuljgwfoevlabixyclzsph
krdmtuojgwfoeknabixyclzsrh
krdmtuqjtwfoevnabiypclzsph
krvmtupjgwfoevnabixycldsph
krdmtuxjgwfoevaabxxyclzsph
krdmtvqlgwfoehnabixyclzsph
wrdmtuqjgwfoevnabixyclzdpy
krdatuqlgwfoevnabixyclzsjh
krdmtuqjgwfoevpabkxyclzsjh
krdmtuqjgwqvsvnabixyclzsph
krdmtwqjgwfoevnobixyclzspm
krdmtuqjgssoevnabixyclgsph
krdmtuqjgwfoevnafixyclzbpp
krdmtuqjowfoevxabiuyclzsph
krdmtuqrgwfoevntbixyclzspu
krdmtucjgwfoevnabixcnlzsph
krddtuojgwfoevnabixyclzzph
krdmtuqjgwuoevnabiqycldsph
kpdmpuqjgwfoevnabixyclzslh
krdmtuqjgwfoewnabixyzxzsph
krdmtuejswfoevhabixyclzsph
krdmtuqggwfoevntbikyclzsph
krdmtuqjgwfoevnabixydlhnph
krdmtcqjglfoevnaxixyclzsph
krumyuqjgwfoevnrbixyclzsph
kgdmmuqjgwooevnabixyclzsph
krdmteqjgwfqevwabixyclzsph
krdmfuqjgwfpevnabixyclzspq
erdmtycjgwfoevnabixyclzsph
krdmcuqjgwfoevnabixjglzsph
krdmtuqjgtfoeunabixiclzsph
krdmtuqjgwfoevmqbixyclzspu
krlmtuqjvwfoevnabikyclzsph
krdotuqjgwfoevnagrxyclzsph
krdmtuqbgwfoefnabixyclasph
kwdmtuqjgwfosjnabixyclzsph
kydmtuqjgwfoevcabixycezsph
crdmtuqjgwloevnabixkclzsph
krimtuqhgwfoevnbbixyclzsph
krdmjuqagwfoevnabicyclzsph
krdmtuqdgzfoevnabixydlzsph
krdmtuqjgwwoevnaqixyclzspf
krdmtuqjgwfoevnabdxyzvzsph
krdmtuqjgwaofvnabixyclzsnh
krdmturjgwfmevnabixyclzspn
krdmvuqjgwboevnabixyolzsph
krdmtuqjgwfomvnabijyclzspx
bedmtuqjgwfoevnabixyslzsph
krdmtenjgwfoevnabixyclzsqh
krdmtuqugwfoevnabixpcdzsph
krdmtuqjgiloevnabrxyclzsph
krdmtupjcwfoevnabixyclwsph
kremtuqjgwfoevnabixyyjzsph
krdmtuqjgwnoovnabixyclzshh
qrdmtuqjgwfoevnabixyciasph
krdituqjgbfoevnagixyclzsph
krdmnoqjgwfoqvnabixyclzsph
krdmtuqegwfoevhkbixyclzsph
krdmkucjgwfoevnabixnclzsph
krdmtuqbnwpoevnabixyclzsph
krdmttqjgwfoevnabixyclbspz
srdmtubjgwfrevnabixyclzsph
krdmruqjzwfoevnabixyclesph
ardmtuqfgwwoevnabixyclzsph
yrumtuqjgwhoevnabixyclzsph
rrdmtuqjgwfoevnabsxycwzsph
krpmtuqjgwfoevdabixyclzzph
krdmuuqjgwfoevnabixyclriph
krdmtuqjgwfobvnabixyvgzsph
krdmbuujgwfoevnabixycczsph
krhmtuwjgwfoeqnabixyclzsph
krdwtuqjgwfoevnkbixyclzzph
krdmtuqjgwkoeqnabixyvlzsph
kadmtuqjgwfoednabcxyclzsph
krdmtyqqgwfoevnabizyclzsph
krdmtuqjgnfoevnabiyycmzsph
krdmtuqjcwfouvnabixyclznph
krdmtuqjjwfcevnqbixyclzsph
krdmtuqfgbfoevgabixyclzsph
kkdmtuqjgwfoevnapixyclzsth
nrdmtuqjgwtoevnakixyclzsph
krdmtuqjglfoevlabixdclzsph
zrdmtuqjgwfoevndbixbclzsph
krdmeuqjgwfoeenabixyclrsph
krdmoaqjzwfoevnabixyclzsph
krsmtuqjgwfoevnwbixyclzsfh
kadmtuqjgwfoqdnabixyclzsph
krsmtuqjgofoevnabixkclzsph
krdmtuqjdwfoevnibixdclzsph
mrdmtuqjgwfouvnabixyclzfph
trdmtlqjgwfoevnabixyclzjph
trdmyuqjgwfozvnabixyclzsph
krdmtiqjgwroevnabixyclzspk
erdmtutjgwftevnabixyclzsph
krdwyuqjgwfoevnaaixyclzsph
krdmthqbgwfoevnabixyclksph
krdmttqjgwfoivnabixyclvsph
krdmtuqjgwfoefnabixyflgsph
khdmtuqjgwfoevnajixyvlzsph
krdmtuqvgwfoevnasixyclzspt
krdmtuqjgkwogvnabixyclzsph
krdmtuqjgwfoevnaboxpglzjph
kadmtuqjgwfoxvnabixyclziph
krdmtuqjfwfoevnabaxycbzsph
krdjtuqjgwfoevnabiryhlzsph
krdvtuqjgpfoevnabcxyclzsph
brdmtuqjgwfoevnafixyqlzsph
krdmtuqjgwfoevnavixxcllsph
krdhtuqjkwfoevfabixyclzsph
krdmtuqjgjfoevnawixyclzsuh
krddtuqjgwfoeqnabiwyclzsph
krhmtuqjgwfnevnabinyclzsph
kedmtuqjgzfmevnabixyclzsph
qrdmtuqjgwfoevntbixyclzxph
krdmtuqsgwfoevnabixvclzrph
scdmtuqjgwfoevnabixtclzsph
krymtuqjgjfolvnabixyclzsph
krdmtuqjgwfkevnablxyclzskh
krymtuqjswfoevnabixyclzvph
krdmtuqjhwfoevnabixycwzspd
krdmtuxjgwfoevnabyxyclzzph
krdmtlqjgwfovvnabilyclzsph
krdmtuqjgwfoevnaaijcclzsph
krdatrqjgwfokvnabixyclzsph
krdmtuqjgwfoevnaxifyclzkph
krddtuqjgwfoevnabixccozsph
krdmtuqngwfoevnabiyycxzsph
krdmtumdgwfoevnqbixyclzsph
krdmtuqjgwfoevnabixyxlmsch
krdmtudzgwfoevnabixtclzsph
krdmtuqjgwfoevnpbixyclhspl
krdmtqqjgwjoevnabexyclzsph
kydmtuqzgwfoevnabixyclwsph
krdmeucjgwqoevnabixyclzsph
krdmtuqjghfoevjabixyclzspp
krdmtuqjgjfwevnabixyclzskh
krdmkuhjgwfoevnabipyclzsph
krdytuqjgwfoevnabibyclztph
krdmtuqjgwfpevnabisyzlzsph
kmdmtgqjgwfsevnabixyclzsph
krdmtuqjgsfoevnabijyclzszh
krdmtuqjgwfoevnabivyclzuuh
krdstuqjgrfoevnabixyclzspu
jrdmtuqjgwfotvnabixyclzspj
krdmrumjgwfoevnabixeclzsph
krpmtusjgwfoevnabioyclzsph`
