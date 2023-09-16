// Made by Fabián José Fernández Fernández

let nth (index: int) (list: int list) : int option =
    let indexes = [0 .. List.length list - 1]
    let indexedList = List.zip indexes list
    List.tryFind (fun (i, _) -> i = index) indexedList
    |> Option.map snd

let exampleList = [1; 2; 3; 4; 5]
let result1 = nth 2 exampleList
let result2 = nth 3 exampleList
let result3 = nth 6 exampleList

printfn "nth 2: %A" result1
printfn "nth 3: %A" result2
printfn "nth 6: %A" result3