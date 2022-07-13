package main

import "fmt"

func main() {

    ex_jnpr := make([]string, 5)
    ex_jnpr[0] = "Devang"
    ex_jnpr[1] = "Mukesh"
    ex_jnpr[2] = "Mayank"
    ex_jnpr[3] = "Vishal"
    ex_jnpr[4] = "Vinay"
    fmt.Println("Ex JNPR List:", ex_jnpr)

    g_jobs := make(map[int]string)
    g_jobs[0] = "GEN"
    g_jobs[1] = "EDGE"
    g_jobs[2] = "ICN"
    fmt.Println("Google Teams:",g_jobs)

    g_db := make(map[string][]string)
    g_db["GEN"] = []string{"Mukesh", "Mayank", "Vishal"} 
    g_db["EDGE"] = []string{"Devang"}
    g_db["ICN"] = []string{"Vinay"}
    fmt.Println("New Team Mapping:", g_db)

}

