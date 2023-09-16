module RutaCorta

open Recipientes

// Grafo de prueba con pesos
let grafoConPesos = [
    ("i", "a", 3);
    ("i", "b", 6);
    ("a", "c", 5);
    ("a", "d", 2);
    ("b", "c", 1);
    ("b", "d", 4);
    ("c", "x", 2);
    ("d", "f", 3);
    ("f", "d", 1);
    ("x", "c", 2)
]

// Función para generar vecinos con pesos
let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | (origen, destino, peso) :: rest ->
        if origen = nodo then (destino, peso) :: vecinos nodo rest
        else vecinos nodo rest

// Función para extender una ruta con pesos
let extender ruta grafo = 
    List.filter (fun x -> x <> [])
        (List.map (fun (nodo, peso) -> 
            if (miembro nodo ruta) then [] 
            else (nodo, peso) :: ruta) (vecinos (List.head ruta) grafo))

// Función principal de búsqueda en profundidad con pesos
let rec prof2 ini fin grafo =
    let rec prof_aux fin ruta peso grafo =
        if List.isEmpty ruta then (([], System.Int32.MaxValue), System.Int32.MaxValue)
        elif (List.head (List.head ruta)) = fin then
            ((List.rev (List.head ruta), peso), peso)
        else
            let vecinosConPesos = vecinos (List.head ruta) grafo
            let rutasExtendidas =
                vecinosConPesos
                |> List.filter (fun (nodo, _) -> not (miembro nodo ruta))
                |> List.map (fun (nodo, pesoVecino) ->
                    let rutaExtendida = (nodo, pesoVecino) :: ruta
                    let pesoExtendido = peso + pesoVecino
                    prof_aux fin rutaExtendida pesoExtendido grafo)
            List.minBy (fun ((_, p), _) -> p) rutasExtendidas
    prof_aux fin [(ini, 0)] 0 grafo





        
    
