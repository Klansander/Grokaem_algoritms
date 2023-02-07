package main



func search_key(graph []string,flag *map[string]bool, graph_full map[string][]string,start_key string,res_key string,result *int) {
	value:=[]string{}
	*result++
	for _,val:= range graph {
		if val==start_key {
			continue
		} else if !(*flag)[val] {
			if val!=res_key {
				value=append(value,graph_full[val]...)
			}else {
				return
			}
			(*flag)[val]=true
		}

	}
	search_key(value,flag,graph_full,start_key,res_key,result)

}
func runAlgo( graph map[string][]string, start_key,res_key string) (result int) {
	flag := map[string]bool{}
	search_key(graph[start_key],&flag,graph,start_key,res_key,&result)
	return
}

func main() {
	graph:= map[string][]string {
		"a": {"b","d"},
		"b": {"c","d"},
		"c": {"b","f"},
		"d": {"b","i"},
		"f": {"c","j"},
		"i": {"j","d"},
		"j": {"i","f"},
		
	}
	println(runAlgo(graph,"a","j"))
}