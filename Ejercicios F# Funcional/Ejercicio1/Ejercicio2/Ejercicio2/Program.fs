// Made by Fabián José Fernández Fernández

let filterBySubstring (substring: string) (strings: string list) : string list =
    strings
    |> List.filter (fun str -> str.Contains(substring))

let exampleStrings = ["la casa"; "el perro"; "pintando la cerca"]
let filteredStrings = filterBySubstring "la" exampleStrings

printfn "Filtered strings: %A" filteredStrings
