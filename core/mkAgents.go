// Copyright 2018 The go-Dacchain Authors
// This file is part of the go-Dacchain library.
//
// The go-Dacchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-Dacchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-Dacchain library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"github.com/Dacchain/go-Dacchain/core/types"
	"strconv"
	"github.com/Dacchain/go-Dacchain/rlp"
)

type genesisAgents []types.Candidate


const (
	firstDelegateVoteNumber  = uint64(1000000000)
	secondDelegateVoteNumber = uint64(300000000)
	thirdDelegateVoteNumber  = uint64(100000000)
	fourthDelegateVoteNumber = uint64(10000000)
	fifthDelegateVoteNumber  = uint64(5000000)
	sixthDelegateVoteNumber  = uint64(1000000)
)

func main() {
	var list genesisAgents
	candidateList := mainNetAgents()
	list = append(list, candidateList...)

	data, err := rlp.EncodeToBytes(list)
	if err != nil {
		panic(err)
	}
	result := strconv.QuoteToASCII(string(data))
	fmt.Println("const agentData =", result)
}

func mainNetAgents() []types.Candidate {
	return []types.Candidate{
		{"0x2da92c7f21c5e7f67c45c0bcf49b3815791bddb4", firstDelegateVoteNumber, "Davinci", 1526095302},
		{"0xd119b77199078028051fee8a86222e0457a8624e", firstDelegateVoteNumber, "Zeus", 1526095302},
		{"0xb75c76088c82e4b47d471fb1097ef96dca48fe01", firstDelegateVoteNumber, "Hera", 1526095302},
		{"0x764f65e0c126eda1499eb28e96db8a318e99d3c0", firstDelegateVoteNumber, "Poseidon", 1526095302},
		{"0x8f9344230671730b97d0a21237ccf4ca4c277d7a", firstDelegateVoteNumber, "Hades", 1526095302},
		{"0x3f5c04d3c579132b15955b4db24e93add96a6645", firstDelegateVoteNumber, "Demeter", 1526095302},
		{"0x75b63c1c2dd493ebfae710056158c1e6a5dc1a21", firstDelegateVoteNumber, "Ares", 1526095302},
		{"0xd51a0f5bab3830a5b07e153b06eef6942ce0d9b6", firstDelegateVoteNumber, "Athene", 1526095302},
		{"0x00bfad3006f42c801e8eeb412442f374a792c9dc", firstDelegateVoteNumber, "Apollo", 1526095302},
		{"0xb7e763bf7822d06cab7e07d1235ea162fd4bd52d", firstDelegateVoteNumber, "Artemis", 1526095302},

		{"0x10366172a384e1f687329af19e018b1ae1b7420c", secondDelegateVoteNumber, "Aphrodite", 1526095302},
		{"0x683ae0b79eb94e41ce32a454f302b08d1a7ef6b2", secondDelegateVoteNumber, "Hermes", 1526095302},
		{"0x37b76f3d56a3ef0149e61bdc9387769e0dc62652", secondDelegateVoteNumber, "Hephaestus", 1526095302},
		{"0x6d724089995dc079938b11313562f522633ef321", secondDelegateVoteNumber, "Andromeda", 1526095302},
		{"0x37314d4179165eb66c65e8f48d6d5ee9ba45b827", secondDelegateVoteNumber, "Antlia", 1526095302},
		{"0x64b7f62e346b5d1b7c70fdfc49b7903a633e7aca", secondDelegateVoteNumber, "Apus", 1526095302},
		{"0x02cdf6a974f9a50ac83b981d5513325eab7ad54c", secondDelegateVoteNumber, "Aquarius", 1526095302},
		{"0x9d779e74e6b662fa9ec86b63e98e62f2cba076d1", secondDelegateVoteNumber, "Aquila", 1526095302},
		{"0x5c280157e1d4efc98c0cbf91cfb56c5b9f38f796", secondDelegateVoteNumber, "Ara", 1526095302},
		{"0x9a1c2d9c4530e634e85f7163413640cd82aba43f", secondDelegateVoteNumber, "Aries", 1526095302},
		{"0x3f922cf5d5566d99af65585b5882e9819ac08f37", secondDelegateVoteNumber, "Auriga", 1526095302},
		{"0x25f919c0a9a918f7e21c459d7f975b7fd266755a", secondDelegateVoteNumber, "Bootes", 1526095302},
		{"0x19735972e6062fc00aefd570dd12ba65d4cff310", secondDelegateVoteNumber, "Caelum", 1526095302},
		{"0x532e4a20be903e82a23a99245d2cea5ce3ed81aa", secondDelegateVoteNumber, "Camelopardalis", 1526095302},
		{"0xf97106cc0385d6a768455df4c23a76100fc6c19a", secondDelegateVoteNumber, "Cancer", 1526095302},
		{"0xe5cc7b02e66f265b2bb77c080b6aa9d850c59fb9", secondDelegateVoteNumber, "Canes Venatici", 1526095302},
		{"0x3d18f900b2a341e2ad5bedbc824927521956c7a0", secondDelegateVoteNumber, "Canis Major", 1526095302},
		{"0x9e47756de4410b06519a9218c954688310e525de", secondDelegateVoteNumber, "Canis Minor", 1526095302},
		{"0xb598d7e3b335ce42310004150c073bd78b289e58", secondDelegateVoteNumber, "Capricorn", 1526095302},
		{"0x511391688396bbb4c4079a89d8fcc800f5b8826c", secondDelegateVoteNumber, "Carina", 1526095302},

		{"0xe8a242f24676bf98044c8ad7259f19597c5c5095", thirdDelegateVoteNumber, "Cassiopeia", 1526095302},
		{"0xfc2d5f1178ddcfd9f1dc76d5c7341b40d0798c98", thirdDelegateVoteNumber, "Centaurus", 1526095302},
		{"0xc475cfb006601871b631b064db148a01db78cb3f", thirdDelegateVoteNumber, "Cepheus", 1526095302},
		{"0xe4462e7890c04643f9758cb777f7907853f621a8", thirdDelegateVoteNumber, "Cetus", 1526095302},
		{"0x8fbd218d2c2b588f5783988695745bedaff31cf2", thirdDelegateVoteNumber, "Chamaeleon", 1526095302},
		{"0xc56f350ad64535f14874db8fef2e9a66eea48e32", thirdDelegateVoteNumber, "Circinus", 1526095302},
		{"0x7019da25e1db7f8f28ac5b4be877cc805e245b9a", thirdDelegateVoteNumber, "Columba", 1526095302},
		{"0x898e34dda831b96e7c49cdd00a5d4457cfc70486", thirdDelegateVoteNumber, "Coma", 1526095302},
		{"0xd4f77c00ccf1b08c4d846fcc37158eb54a42845e", thirdDelegateVoteNumber, "Corona Australis", 1526095302},
		{"0x0d1afbde57cccab42217c4a4309763fb54c58e90", thirdDelegateVoteNumber, "Corona Borealis", 1526095302},
		{"0x945509c5cb753b06ab52a2b688ed1a03ab7efe6d", thirdDelegateVoteNumber, "Corvus", 1526095302},
		{"0xd26b8ac25221a952724ebcec5ddbc36f91b8570a", thirdDelegateVoteNumber, "Crater", 1526095302},
		{"0x0192b1116a9574465503cea57f28970139c326b7", thirdDelegateVoteNumber, "Crux", 1526095302},
		{"0xbb670374c3e2450ddb3d931af6fa999d8a3d3525", thirdDelegateVoteNumber, "Cygnus", 1526095302},
		{"0xca77a6302079d0649658b8613730f0fa00c40feb", thirdDelegateVoteNumber, "Delphinus", 1526095302},
		{"0x5a58c1b1027a7d9af7244697e145368c31e49232", thirdDelegateVoteNumber, "Dorado", 1526095302},
		{"0x3c5fd290570f69f6046ba9278547372efdac7a11", thirdDelegateVoteNumber, "Draco", 1526095302},
		{"0x6fb1cd24b42fd89b3676d3d996cf211a82ccf2c6", thirdDelegateVoteNumber, "Equuleus", 1526095302},
		{"0xaf3f96fda58ca9450695932e5b23bf45738d8238", thirdDelegateVoteNumber, "Eridanus", 1526095302},
		{"0x4278c4ffe00c565a9993a11ba54ece60330cd6a4", thirdDelegateVoteNumber, "Fornax", 1526095302},
		{"0x30b4fe5320a28c89c59004370c08331e3681b295", thirdDelegateVoteNumber, "Gemini", 1526095302},
		{"0x89ec110be932c98651734a1c6617b6705681d78f", thirdDelegateVoteNumber, "Grus", 1526095302},
		{"0x9154d12ff2f69a7ff617d08117516e4f191deba0", thirdDelegateVoteNumber, "Hercules", 1526095302},
		{"0xda20e2abfa7742971baff262bf32679b665a89c1", thirdDelegateVoteNumber, "Horologium", 1526095302},
		{"0x79ee9c38be44c831f495f4968b8e1d3107e03bc2", thirdDelegateVoteNumber, "Hydra", 1526095302},
		{"0xc8a1bc1fc8ae598487a819e1cadafa82177594f5", thirdDelegateVoteNumber, "Hydrus", 1526095302},
		{"0x486f0db85e21fd28da12449622923f719e11a052", thirdDelegateVoteNumber, "Indus", 1526095302},
		{"0x865d50f3077dd2e071c61d967e744d8514c27aef", thirdDelegateVoteNumber, "Lacerta", 1526095302},
		{"0xc15a359d3fd3d7156d6bfd162895b08d5a896d48", thirdDelegateVoteNumber, "Leo", 1526095302},
		{"0x7502162773c8e5e72913b7b5ffc4f041a580574e", thirdDelegateVoteNumber, "Leo Minor", 1526095302},
		{"0xdab320bc61bab9c1914f07bf581d827abd12f256", thirdDelegateVoteNumber, "Lepus", 1526095302},
		{"0x66943a1af837a658144d4b7866984fd081c89c2e", thirdDelegateVoteNumber, "Libra", 1526095302},
		{"0x697b7b80d2b67fa6594a33411546c06a62e3a553", thirdDelegateVoteNumber, "Lupus", 1526095302},
		{"0xef85dc8036c397353e8f4c5ca32effc7c7628d71", thirdDelegateVoteNumber, "Lynx", 1526095302},
		{"0xce8d1a0de093b16d98acafd573ec40d6cead1c5a", thirdDelegateVoteNumber, "Lyra", 1526095302},

		{"0xaff28b74610ad3007bb954ff6189aabdc7fdbd34", fourthDelegateVoteNumber, "Mensa", 1526095302},
		{"0xf2ca1e7b0b49ecba7b8e5ac14704c06f6295cb6f", fourthDelegateVoteNumber, "Microscopium", 1526095302},
		{"0x5074f84c66e7882bd95523f5ff6a0b79b0b1f171", fourthDelegateVoteNumber, "Monoceros", 1526095302},
		{"0xdcbf1cc70ea571a7246b0782a182255c23159657", fourthDelegateVoteNumber, "Musca", 1526095302},
		{"0x2ae5c685546b50746a740edb23c66d6f0796cf88", fourthDelegateVoteNumber, "Norma", 1526095302},
		{"0xfff2d75669eb6ed6978220d4d627cf296dc13ea6", fourthDelegateVoteNumber, "Octans", 1526095302},
		{"0xfcfe4ae31c022162fb77a39f8f0af99795eb9857", fourthDelegateVoteNumber, "Ophiuchus", 1526095302},
		{"0xad1d73e4915b74b962f910aa15ac123dee63a090", fourthDelegateVoteNumber, "Orion", 1526095302},
		{"0x431cc714b9bd2be25f902538aafa4e341697847f", fourthDelegateVoteNumber, "Pavo", 1526095302},
		{"0x82005c0922aa243d287185baba530c42bdf3a187", fourthDelegateVoteNumber, "Pegasus", 1526095302},
		{"0xe8ff4b6ccdc8bd36e6f94b48a64be794cad1d273", fourthDelegateVoteNumber, "Perseus", 1526095302},
		{"0xc5fc0377d874539758e2caa1b150506ccff46530", fourthDelegateVoteNumber, "Phoenix", 1526095302},
		{"0x67b4f0cfbe7065e9e3a3ca81a1f1d28bb5442f72", fourthDelegateVoteNumber, "Pictor", 1526095302},
		{"0x0321bb78a0fe12450c7bf38c637a5b94fe64128f", fourthDelegateVoteNumber, "Pisces", 1526095302},
		{"0x9f2a0ddc022bc5fda8d47c565e388bccd8df38ca", fourthDelegateVoteNumber, "Piscis Austrinus", 1526095302},
		{"0x20e0afe3b2ac112217cb7c2259e12a4d8adb9a58", fourthDelegateVoteNumber, "Puppis", 1526095302},
		{"0x3a66d18be0f811f28212ba43a41507213f32b1d0", fourthDelegateVoteNumber, "Pyxis", 1526095302},
		{"0xa04ea8ebf1c73a9a0a2457ee7746545bafd95d0b", fourthDelegateVoteNumber, "Reticulum", 1526095302},
		{"0xe2a59245f595205323208ffcbb0a7329afddb28a", fourthDelegateVoteNumber, "Sagitta", 1526095302},
		{"0x2b249d8743fcef9740d52055f181f7e3f933387b", fourthDelegateVoteNumber, "Sagittarius", 1526095302},
		{"0x3fd66fb30577ab8a2a3d973f8e92c0f83e420b32", fourthDelegateVoteNumber, "Scorpius", 1526095302},
		{"0x052a1e1fae1586b0a93f23dc761255b16692ec1d", fourthDelegateVoteNumber, "Sculptor", 1526095302},
		{"0xfc5462e5b47253b10a8cef9ce8e0d55388ab72d0", fourthDelegateVoteNumber, "Scutum", 1526095302},
		{"0x2fb9e102254831c382fe53c508659930c31ba574", fourthDelegateVoteNumber, "Serpens", 1526095302},
		{"0xb0f74ecb4b087722909bca017e6b3c4405c52b7d", fourthDelegateVoteNumber, "Sextans", 1526095302},
		{"0x3108d42b1744fd997f6fc045d12153505d634e20", fourthDelegateVoteNumber, "Taurus", 1526095302},
		{"0x2727fe0b8be0f3aaacafcdc547aff1d2aff3889f", fourthDelegateVoteNumber, "Telescopium", 1526095302},
		{"0x4bd48ef4f90fd3ec5a57ce8925839614f9d2331f", fourthDelegateVoteNumber, "Triangulum", 1526095302},
		{"0xc91be4f81387d8acabb477d88e824ce84d3fb0fd", fourthDelegateVoteNumber, "Triangulum Australe", 1526095302},
		{"0x1422b8d66cfa693a2cd088545e7436c8bf7bf051", fourthDelegateVoteNumber, "Tucana", 1526095302},

		{"0x9c0fe219f448d18f2bb4bfbeaf320a3893fbbdbb", fifthDelegateVoteNumber, "Ursa Major", 1526095302},
		{"0x800f30a1e883583235f1f597cdb47c1f5f0da62c", fifthDelegateVoteNumber, "Ursa Minor", 1526095302},
		{"0xd20cf6e9b1afd6c64f43abaac7ffe0c76d79327b", fifthDelegateVoteNumber, "Vela", 1526095302},
		{"0xb4e944dd6f37d0628478fa6418cbe69b12f67dfa", fifthDelegateVoteNumber, "Virgo", 1526095302},
		{"0xd8f6717f5f2f9f904436089affc0599c10323c75", fifthDelegateVoteNumber, "Volans", 1526095302},
		{"0xeaa42c139bb737d38aba79be40003c8b9b42d5a0", fifthDelegateVoteNumber, "Vulpecula", 1526095302},

		{"0x8da41f942533b4835a7edb9d85faa79ac765e9f8", sixthDelegateVoteNumber, "Vulpecula", 1526095302},
	}
}

func mainTestNetAgents() []types.Candidate {
	candidateList := []types.Candidate{
		{"0x9d797788c3a3adace9155cfead1c35224e49fa15", uint64(2000000), "dac-node1", 1492009146},
		{"0x0315141103b566e3ac49d0011636f88bd1f560dc", uint64(2000000), "dac-node2", 1492009146},
		{"0x6c465c2923ee22d160c925e8b3ef91b202d94445", uint64(2000000), "dac-node3", 1492009146},
		{"0x22c2d6381213406269452761c15c354190dbc687", uint64(2000000), "dac-node4", 1492009146},
		{"0x55620e3b0d8ecba7039d8b8ef7a5d6c28e6a6d51", uint64(2000000), "dac-node5", 1492009146},

		{"0xf104acc210245428e0e515b2e44328e821d043e7", uint64(2000000), "dac-node6", 1492009146},
		{"0xa42f95777521ddedb8629e8f473655cbb583f949", uint64(2000000), "dac-node7", 1492009146},
		{"0xa6a2f80cc864b546d8bb8d3505b24828d41e3d3d", uint64(2000000), "dac-node8", 1492009146},
		{"0x047703ea5a2ff7671344acbbda87356a59503ab3", uint64(2000000), "dac-node9", 1492009146},
		{"0xd7325af63f016165dc2b0696924af19d04fe29e8", uint64(2000000), "dac-node10", 1492009146},


		{"0xd7519dfc1cf9f8df2b5fa50c3536ed7c762216d3", uint64(2000000), "dac-node11", 1492009146},
		{"0x12aa15390f5cd0fa308f08a3f18253455ee163d3", uint64(2000000), "dac-node12", 1492009146},
		{"0x6c6ac8ed31d9eae67b9e5929ade9bc8d232d124d", uint64(2000000), "dac-node13", 1492009146},
		{"0x99b395532231b7ae3cc953ff7550ab5beae2724d", uint64(2000000), "dac-node14", 1492009146},
		{"0x606894726b4d1bbe7c6e28221bf47ce56d20e599", uint64(2000000), "dac-node15", 1492009146},


		{"0x72554e84e755b5813cf6ec8341fd9a80c7bc602a", uint64(2000000), "dac-node16", 1492009146},
		{"0xe96f6ae493e4cf1342de29c23e2e3d885a41a042", uint64(2000000), "dac-node17", 1492009146},
		{"0x432a324d1bc0407f90c982f68373ee10a60a2eac", uint64(2000000), "dac-node18", 1492009146},
		{"0x97030507ae987fb11ee3e3643db93b7f01be77cc", uint64(2000000), "dac-node19", 1492009146},
		{"0x42d235713441b6277e57edcbd24b0a2b553bccea", uint64(2000000), "dac-node20", 1492009146},


		{"0x6f8c270c024af65fbd4f2c8235c173b64d2d7a05", uint64(2000000), "dac-node21", 1492009146},
		{"0xae8cc2a8e4c560ac680098b1a2af9a19cafe9a40", uint64(2000000), "dac-node22", 1492009146},
		{"0x8a9bf1c4a6a9a294562fa42bdc1a42c90e8dfd3e", uint64(2000000), "dac-node23", 1492009146},
		{"0x1520afd75ed3236687096acbf872634826e33cce", uint64(2000000), "dac-node24", 1492009146},
		{"0x5633852a7daf338eb6e75babc7b2c3a0511f8d34", uint64(2000000), "dac-node25", 1492009146},


		{"0x9f5b2287d2c668792abe10e181a24463c7a9da8a", uint64(2000000), "dac-node26", 1492009146},
		{"0xd249122237fad72e27590e5eacbb9d96d2e9e3d5", uint64(2000000), "dac-node27", 1492009146},
		{"0x881b48171f25fe005593caa5d75d5e51bdfc02e5", uint64(2000000), "dac-node28", 1492009146},
		{"0x7ec08ba4f63a69b756fd78080945a6e323258470", uint64(2000000), "dac-node29", 1492009146},
		{"0x258b9f3351f6f9a99c88e02ff4ae8226d120d74c", uint64(2000000), "dac-node30", 1492009146},


		{"0x9679e8d7ded06a2039dcb197da4a090506ce82d2", uint64(2000000), "dac-node31", 1492009146},
		{"0x90927168e1c192d74068066ffeb0d723c8208d56", uint64(2000000), "dac-node32", 1492009146},
		{"0x4f1ceb3e0294a1a7ff5148974c2e47943ae03343", uint64(2000000), "dac-node33", 1492009146},
		{"0x458a71bc2150d590f08a8139dfb39bc53536169c", uint64(2000000), "dac-node34", 1492009146},
		{"0x3693d4df12a30e9a39d7b80dd5601c3e04cd2beb", uint64(2000000), "dac-node35", 1492009146},


		{"0xdd9b8aa04ec9ecf09b68913ec4ec3afdd05882fb", uint64(2000000), "dac-node36", 1492009146},
		{"0xa00e6742526ae22759fd2e02ac3ecbe90f37af75", uint64(2000000), "dac-node37", 1492009146},
		{"0x2da04c6f71e64b57eeb06bf24183f354426d340e", uint64(2000000), "dac-node38", 1492009146},
		{"0xd1bd9f32b6e4b821ce2a3ae68c2626788d864365", uint64(2000000), "dac-node39", 1492009146},
		{"0x9bdfb96dd30622f1b8f5cf27ba9e5445ed5b7400", uint64(2000000), "dac-node40", 1492009146},


		{"0x7d2a545617a431874235f4e098ed81299e3f5abc", uint64(2000000), "dac-node41", 1492009146},
		{"0x8cf0269ad4ba69c364f9c4831339e232dcee5a86", uint64(2000000), "dac-node42", 1492009146},
		{"0x26ac91ef5d627c528375bb2ddbceb7aef7f07081", uint64(2000000), "dac-node43", 1492009146},
		{"0x23e31583f69beeace957b8d100ac328763604fb3", uint64(2000000), "dac-node44", 1492009146},
		{"0x0a2d7b2ddb29b24ef647533b8dffd4daeca877f1", uint64(2000000), "dac-node45", 1492009146},


		{"0x66e532846a8068cfac2dcd5b0fa950d887cf83e2", uint64(2000000), "dac-node46", 1492009146},
		{"0x73773efe6ff6284c36ea0013cf3978c432d19405", uint64(2000000), "dac-node47", 1492009146},
		{"0x5d323e4a3d1ac1ee71fbf07c9568de28a9fdfbf3", uint64(2000000), "dac-node48", 1492009146},
		{"0x396d1700f9a0c5e8ba34813bc74c312fc988cca1", uint64(2000000), "dac-node49", 1492009146},
		{"0x7c446ceb9ce416ba20e3a64a3493a13873453903", uint64(2000000), "dac-node50", 1492009146},


		{"0xf45e8026147fc47afc78304367ae03bf0ffc127a", uint64(2000000), "dac-node51", 1492009146},
		{"0x37f9ff4f60e428ffaa8e84a5dfbfa70422ca45ea", uint64(2000000), "dac-node52", 1492009146},
		{"0x6cb00266e649ac26c2347161d572bafd285719fd", uint64(2000000), "dac-node53", 1492009146},
		{"0x3645e4d275a4d36c4b43c683f8868b13fd595e93", uint64(2000000), "dac-node54", 1492009146},
		{"0x49defe4396ced1345e960cdbb20cabdd02eeeb32", uint64(2000000), "dac-node55", 1492009146},


		{"0x4c694deded8f660a16ecf22030d4c61f9b0f7621", uint64(2000000), "dac-node56", 1492009146},
		{"0x3509a4e1337e9c21d745465bede8535e3b95c5f7", uint64(2000000), "dac-node57", 1492009146},
		{"0x5b5214eb06f3ee53470d1a374a2d62a1b9a07036", uint64(2000000), "dac-node58", 1492009146},
		{"0x7657138e686299bdac21ec9c9d35935bdbcfe66d", uint64(2000000), "dac-node59", 1492009146},
		{"0xf7eeff70d6a23e2fd6b0054f559ba5ad5920f972", uint64(2000000), "dac-node60", 1492009146},


		{"0xf884e1e6f9e2f500df8d0e18b5857d1e8654acac", uint64(2000000), "dac-node61", 1492009146},
		{"0xd17945e5213730d87619151f6a3ba7b35dd852d6", uint64(2000000), "dac-node62", 1492009146},
		{"0x260a286ed1ad6ba8b917f67b14e2c96b73b33173", uint64(2000000), "dac-node63", 1492009146},
		{"0x88ed077e9a648f4dd51fce59d3c67f02fdc7b06b", uint64(2000000), "dac-node64", 1492009146},
		{"0x8ff287eda0cfd0ecac07dff934bab7324be217ab", uint64(2000000), "dac-node65", 1492009146},


		{"0xada53089ea62d86e06a380de2c0e0e217653f2a7", uint64(2000000), "dac-node66", 1492009146},
		{"0xfeaeb528993b4fd745dffd0e7fad173284fd7955", uint64(2000000), "dac-node67", 1492009146},
		{"0x070cc4553e882d3defafbea7adc9875056faf31d", uint64(2000000), "dac-node68", 1492009146},
		{"0x3306d2bd3e358954707ec064baeb7307f9e5f47b", uint64(2000000), "dac-node69", 1492009146},
		{"0xb4a927838cb05d4160c92d1614debdee6f0243d7", uint64(2000000), "dac-node70", 1492009146},


		{"0xbd32f261b73478e2b717f2e1b21eaf8a8f996d07", uint64(2000000), "dac-node71", 1492009146},
		{"0x72c3573aeff849ae5cda6c15a89c89d5fe9e3d1d", uint64(2000000), "dac-node72", 1492009146},
		{"0x254acca74301e60fecd7b38829211cb3b1520a7e", uint64(2000000), "dac-node73", 1492009146},
		{"0xc23c0f496e38f9eb27c7a544b0317737e1a7d6f6", uint64(2000000), "dac-node74", 1492009146},
		{"0x78b8c2f89c5c2a90f6ef0da1c2081f17e73c2f0f", uint64(2000000), "dac-node75", 1492009146},


		{"0x77ce49313a906f41f3ce1a6ba2631a31eae5d607", uint64(2000000), "dac-node76", 1492009146},
		{"0x9009c5c557a1164481652db12e1c754a81dc88f9", uint64(2000000), "dac-node77", 1492009146},
		{"0x5de0fdca10490b85a7b2327663e1d7f99763b128", uint64(2000000), "dac-node78", 1492009146},
		{"0xdd98dbb2531de5a2bdb96f817465ddda6dee0a06", uint64(2000000), "dac-node79", 1492009146},
		{"0xf13111a1bc2cfb81e7768406c8caeaa6d27b49d7", uint64(2000000), "dac-node80", 1492009146},


		{"0xe2c418d35420a95fe9e64d3431f1cfbcd13843c7", uint64(2000000), "dac-node81", 1492009146},
		{"0x2af84d052436a0bb52b65c98aa309b77675e8088", uint64(2000000), "dac-node82", 1492009146},
		{"0x5326f5601f7723d8ceef123022c8c22c8b1da1ef", uint64(2000000), "dac-node83", 1492009146},
		{"0xaaec00273e9105104b7248228855e5d03232f17f", uint64(2000000), "dac-node84", 1492009146},
		{"0x459c6bf66383d599ffdb9df335ef0c337f0fda22", uint64(2000000), "dac-node85", 1492009146},


		{"0xc1394eea7b0f32645f29dc35e5357b9e7e8bc074", uint64(2000000), "dac-node86", 1492009146},
		{"0x83ffcb95cccf05f4c3d685f88f277c22fceec565", uint64(2000000), "dac-node87", 1492009146},
		{"0xcf514829e8e92249fa9d72f3c7158821cbcc435b", uint64(2000000), "dac-node88", 1492009146},
		{"0x1d50cfa9483ae8bfbdfca0d5a96a054544611d17", uint64(2000000), "dac-node89", 1492009146},
		{"0x73f7b10dfe1697f3ed8be86ef65d308711c604fb", uint64(2000000), "dac-node90", 1492009146},


		{"0x1f21bb8d50fcc7156658bb450e5018e84896f807", uint64(2000000), "dac-node91", 1492009146},
		{"0xed54bc0c55e711ecda1e19d7e31295294d0cf87d", uint64(2000000), "dac-node92", 1492009146},
		{"0xb30bd16fb3de309e6f6d6dfba59a1f50a315ecc3", uint64(2000000), "dac-node93", 1492009146},
		{"0x16be2a356331205b26748912f46f96023a604d9e", uint64(2000000), "dac-node94", 1492009146},
		{"0x7add98d5403b48c3c9c00ebddced2bfd5488123b", uint64(2000000), "dac-node95", 1492009146},


		{"0x6273fd6a5b8fff875a0f3b1363569a3b3d9949ac", uint64(2000000), "dac-node96", 1492009146},
		{"0x2f86611c448c191c1712c36390744ba48ccdf868", uint64(2000000), "dac-node97", 1492009146},
		{"0x46117e313d87f6b005fdf5432957535392122dc5", uint64(2000000), "dac-node98", 1492009146},
		{"0x69e73a911b7ebed06f70f666fe9fcb37980468e8", uint64(2000000), "dac-node99", 1492009146},
		{"0x1928651ca6693b6393524c4f942c686defb8ba31", uint64(2000000), "dac-node100", 1492009146},

		{"0x0cbd0e5cf2508757afb71981bc6f70f1cb58a2aa", uint64(2000000), "dac-node101", 1492009146},
	}
	return candidateList
}
