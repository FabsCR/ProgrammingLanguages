// Made by Fabián José Fernández Fernández

let shiftList (direction: string) (n: int) (lst: int list) : int list =
    let length = List.length lst

    let effectiveShift =
        if length = 0 then 0
        elif direction = "left" then n % length
        else (length - (n % length)) % length

    let shiftLeft (lst: int list) (shiftAmount: int) : int list =
        List.skip shiftAmount lst @ List.take shiftAmount lst

    let shiftRight (lst: int list) (shiftAmount: int) : int list =
        shiftLeft lst (List.length lst - shiftAmount)

    if effectiveShift = 0 then lst
    elif direction = "left" then shiftLeft lst effectiveShift
    else shiftRight lst effectiveShift

let exampleList = [1; 2; 3; 4; 5]
let shiftedLeft = shiftList "left" 3 exampleList
let shiftedRight = shiftList "right" 2 exampleList
let shiftedExcessiveLeft = shiftList "left" 6 exampleList

printfn "Shifted left by 3: %A" shiftedLeft
printfn "Shifted right by 2: %A" shiftedRight
printfn "Shifted excessively left by 6: %A" shiftedExcessiveLeft
