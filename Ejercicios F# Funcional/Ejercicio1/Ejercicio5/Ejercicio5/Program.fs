// Made by Fabián José Fernández Fernández
// This code has modified files that were seen in class, with the teacher and with other classmates.

open RutaCorta
open Recipientes

[<EntryPoint>]
let main _ =
    let ruta = solve con_paredes
    printfn " --> Ruta corta con paredes: %A\n" ruta

    let ruta = solve con_pocas_paredes
    printfn " --> Ruta corta sin las paredes de 8|14 y 25|26: %A\n" ruta

    let ruta = solve_one_path sin_paredes
    printfn " --> Primer camino sin paredes encontrado: %A\n" ruta

    printfn " --> Calculando la mejor ruta sin paredes de todas las posibles ---> ((Esto puede tardar))"
    let ruta = solve sin_paredes
    printfn "Ruta corta sin paredes: %A" ruta

    0b0