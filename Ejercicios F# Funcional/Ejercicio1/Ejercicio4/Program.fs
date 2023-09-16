open Recipientes
open RutaCorta

let pasosRecipientes = prof [2; 2] [2; 0]
printfn "%A" pasosRecipientes

let (rutaCorta, pesoRuta) = prof2 "i" "f" grafoConPesos
printfn "Ruta más corta: %A" rutaCorta
printfn "Peso de la ruta más corta: %d" pesoRuta

