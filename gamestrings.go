package kaa

const gameString1 = `{"you":"0623b12a-411b-4674-a115-591063ef92d3","width":10,"turn":124,"snakes":[{"taunt":"battlesnake-go!","name":"7eef72e9-72fc-4c27-a387-898384639f46 (10x10)","id":"0623b12a-411b-4674-a115-591063ef92d3","health_points":96,"coords":[[9,1],[9,0],[8,0],[8,1],[8,2],[7,2],[7,3],[7,4],[7,5],[7,6],[6,6],[6,7],[5,7],[4,7],[3,7],[2,7],[1,7],[1,8],[0,8],[0,7],[0,6],[1,6],[2,6],[3,6],[4,6],[5,6],[5,5],[5,4],[5,3],[5,2],[6,2],[6,1]]}],"height":10,"game_id":"7eef72e9-72fc-4c27-a387-898384639f46","food":[[0,0],[1,3],[4,0]],"dead_snakes":[]}`

const gameString2 = `{"you":"82557bbc-5ff2-4e51-8133-f6875d4f8d71","width":10,"turn":233,"snakes":[{"taunt":"battlesnake-go!","name":"7eef72e9-72fc-4c27-a387-898384639f46 (10x10)","id":"82557bbc-5ff2-4e51-8133-f6875d4f8d71","health_points":100,"coords":[[1,3],[0,3],[0,4],[0,5],[0,6],[0,7],[1,7],[2,7],[3,7],[3,8],[3,9],[4,9],[4,8],[4,7],[4,6],[4,5],[5,5],[5,4],[6,4],[7,4],[7,3],[6,3],[5,3],[4,3],[4,4],[3,4],[3,3],[3,2],[4,2],[5,2],[5,1],[4,1],[3,1],[2,1],[2,0],[3,0],[4,0],[5,0],[6,0],[7,0],[8,0],[9,0],[9,1],[9,2],[9,2]]}],"height":10,"game_id":"7eef72e9-72fc-4c27-a387-898384639f46","food":[[6,2],[7,5],[2,3]],"dead_snakes":[]}`

const gameString3 = `{"you":"82557bbc-5ff2-4e51-8133-f6875d4f8d71","width":10,"turn":223,"snakes":[{"taunt":"battlesnake-go!","name":"7eef72e9-72fc-4c27-a387-898384639f46 (10x10)","id":"82557bbc-5ff2-4e51-8133-f6875d4f8d71","health_points":100,"coords":[[3,9],[4,9],[4,8],[4,7],[4,6],[4,5],[5,5],[5,4],[6,4],[7,4],[7,3],[6,3],[5,3],[4,3],[4,4],[3,4],[3,3],[3,2],[4,2],[5,2],[5,1],[4,1],[3,1],[2,1],[2,0],[3,0],[4,0],[5,0],[6,0],[7,0],[8,0],[9,0],[9,1],[9,2],[9,3],[9,4],[9,5],[9,6],[9,7],[9,8],[8,8],[8,8]]}],"height":10,"game_id":"7eef72e9-72fc-4c27-a387-898384639f46","food":[[0,5],[0,7],[0,6]],"dead_snakes":[]}`

const gameString4 = `{"you":"0086b0de-4d23-4189-9305-a7308240edb4","width":20,"turn":176,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"3a1cf2b6-ab7f-4870-b672-cb60d7ab4e67","health_points":90,"coords":[[1,15],[2,15],[3,15],[4,15],[4,14],[4,13],[4,12],[4,11],[4,10],[4,9],[3,9],[3,10],[3,11],[2,11],[1,11],[1,10],[1,9],[0,9],[0,8],[1,8],[2,8],[2,7],[3,7],[3,6],[3,5],[4,5],[5,5],[6,5],[7,5],[8,5],[9,5],[9,6],[9,7],[10,7]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"0086b0de-4d23-4189-9305-a7308240edb4","health_points":94,"coords":[[14,16],[13,16],[12,16],[11,16],[10,16],[10,15],[10,14],[11,14],[11,13],[10,13],[9,13],[8,13],[8,14],[8,15],[8,16],[8,17],[8,18],[7,18],[6,18],[6,19],[7,19],[8,19],[9,19],[10,19],[10,18],[11,18],[12,18],[13,18],[14,18],[15,18],[16,18],[17,18],[17,17],[17,16],[17,15],[16,15],[15,15],[15,14],[14,14],[13,14]]}],"height":20,"game_id":"1dd0a217-baef-46ca-bd17-d48a38d54436","food":[[0,6],[16,16],[15,19],[18,7],[8,4],[10,2],[0,15],[3,1],[13,2],[14,11]],"dead_snakes":[]}`

const gameString5 = `{"you":"3de4f206-1538-4bfc-ad49-4cb17fc61bb5","width":10,"turn":8,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"3de4f206-1538-4bfc-ad49-4cb17fc61bb5","health_points":92,"coords":[[5,8],[6,8],[7,8]]}],"height":10,"game_id":"8d58c168-aa8c-4395-90dc-79e36b32bf1e","food":[[2,5],[5,7],[1,8],[8,7],[6,2],[2,8],[7,4],[6,3],[4,2],[6,4]],"dead_snakes":[]}`

const gameString6 = `{"you":"d6f6ff70-fc64-4277-a4a8-b9d0106a87ce","width":20,"turn":376,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"d6f6ff70-fc64-4277-a4a8-b9d0106a87ce","health_points":79,"coords":[[0,13],[0,14],[1,14],[1,15],[0,15],[0,16],[0,17],[0,18],[0,19],[1,19],[1,18],[1,17],[1,16],[2,16],[2,15],[3,15],[3,14],[4,14],[5,14],[6,14],[7,14],[7,15],[6,15],[5,15],[5,16],[5,17],[5,18],[5,19],[6,19],[6,18],[7,18],[8,18],[9,18],[10,18],[10,17],[9,17],[8,17],[8,16],[8,15],[8,14],[8,13],[8,12],[8,11],[7,11],[6,11],[5,11],[4,11],[3,11],[2,11],[1,11],[0,11],[0,10],[0,9],[0,8],[0,7],[0,6],[0,5],[0,4],[0,3],[0,2],[0,1],[1,1],[2,1]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[14,3],[9,19],[17,14],[12,5],[11,6],[12,8]],"dead_snakes":[]}`

const gameString7 = `{"you":"c38c3d40-681c-4f37-8060-bc954e5d1005","width":20,"turn":233,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"c38c3d40-681c-4f37-8060-bc954e5d1005","health_points":96,"coords":[[11,10],[11,9],[11,8],[11,7],[11,6],[12,6],[13,6],[13,5],[13,4],[13,3],[13,2],[13,1],[13,0],[12,0],[11,0],[10,0],[9,0],[9,1],[9,2],[9,3],[9,4],[9,5],[9,6],[9,7],[9,8],[9,9],[9,10],[9,11],[10,11],[11,11],[11,12],[11,13],[11,14],[11,15],[12,15],[13,15]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[6,19],[0,18],[2,13],[0,17],[7,19],[2,1]],"dead_snakes":[]}`

const gameString8 = `{"you":"c08ce082-1ea3-47b1-b347-fd7faa42bdd8","width":20,"turn":606,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"c08ce082-1ea3-47b1-b347-fd7faa42bdd8","health_points":96,"coords":[[16,1],[17,1],[18,1],[19,1],[19,0],[18,0],[17,0],[16,0],[15,0],[14,0],[13,0],[12,0],[11,0],[11,1],[11,2],[11,3],[11,4],[11,5],[11,6],[12,6],[13,6],[13,7],[13,8],[12,8],[11,8],[10,8],[9,8],[8,8],[7,8],[6,8],[6,9],[6,10],[7,10],[7,9],[8,9],[9,9],[10,9],[11,9],[12,9],[13,9],[14,9],[15,9],[16,9],[17,9],[18,9],[18,8],[19,8],[19,9],[19,10],[19,11],[19,12],[18,12],[17,12],[16,12],[15,12],[14,12],[14,11],[14,10],[13,10],[12,10],[12,11],[11,11],[10,11],[9,11],[8,11],[7,11],[7,12],[7,13],[7,14],[8,14],[8,15],[8,16],[7,16],[6,16],[5,16],[4,16],[3,16],[2,16],[2,15],[2,14],[3,14]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[2,11],[2,13],[14,13],[2,10],[7,18],[15,19]],"dead_snakes":[]}
`
const gameString9 = `{"you":"f30c5778-b50b-45a3-8d17-05117ebd35a9","width":20,"turn":125,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"40bae8a0-5d41-4dd4-ae3a-3496eb9a6603","health_points":92,"coords":[[9,2],[10,2],[11,2],[12,2],[12,3],[12,4],[13,4],[14,4],[14,5],[13,5],[12,5],[12,6],[11,6],[11,5],[11,4],[11,3],[10,3],[9,3],[8,3],[7,3],[6,3]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"f30c5778-b50b-45a3-8d17-05117ebd35a9","health_points":89,"coords":[[13,2],[13,3],[14,3],[15,3],[15,4],[15,5],[15,6],[15,7],[14,7],[13,7],[12,7],[11,7],[11,8],[11,9],[11,10],[11,11],[11,12],[11,13],[11,14],[11,15],[11,16],[11,17]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"7a996f7c-2a92-4327-a1d3-464931ae4aa2","health_points":89,"coords":[[13,16],[13,15],[13,14],[13,13],[13,12],[13, 11],[13,10],[13,9],[14,9],[14,8],[15,8],[16,8],[16,9],[16,10],[15,10],[14,10],[14,11],[14,12],[14,13]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[4,18],[5,4],[1,11],[3,16],[0,11],[8,0]],"dead_snakes":[]}
`
const gameString10 = `{"you":"301ac963-6fd0-478b-b1d1-7d9a227b07d4","width":20,"turn":199,"snakes":[{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"0e0d9e0d-5ff7-4cc8-8052-cf49d92bc347","health_points":97,"coords":[[1,6],[1,7],[1,8],[1,9],[0,9],[0,10],[0,11],[0,12],[1,12],[2,12],[3,12],[4,12],[5,12],[6,12],[7,12],[8,12],[9,12],[9,11],[9,10],[9,9],[10,9],[11,9],[11,8],[11,7],[11,6],[10,6],[9,6],[8,6],[8,5],[8,4],[8,3]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"e3e306d5-f5c0-4dc0-9b5e-6896f8ebcd9a","health_points":96,"coords":[[5,19],[4,19],[4,18],[3,18],[3,17],[4,17],[4,16],[5,16],[6,16],[7,16],[8,16],[9,16],[9,15],[9,14],[9,13],[10,13],[10,12],[10,11],[10,10],[11,10]]},{"taunt":"Dad 2.0 Ready","name":"Your New Dad","id":"301ac963-6fd0-478b-b1d1-7d9a227b07d4","health_points":100,"coords":[[7,18],[7,17],[8,17],[9,17],[10,17],[10,16],[11,16],[12,16],[13,16],[14,16],[15,16],[16,16],[17,16],[18,16],[18,15],[17,15],[16,15],[16,14],[16,13],[16,12],[15,12],[14,12],[14,11],[14,11]]}],"height":20,"game_id":"e5fe7eb6-6420-4499-a387-d8537b40c6d1","food":[[15,7],[12,18],[5,6],[18,3],[5,1],[16,2]],"dead_snakes":[]}
`