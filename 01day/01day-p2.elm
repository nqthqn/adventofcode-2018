module Main exposing (main)

import Html exposing (Html, text)
import Set exposing (Set)


ints =
    data
        |> String.split "\n"
        |> List.filterMap String.toInt


main : Html msg
main =
    findSecondSeen ints 0 Set.empty
        |> Debug.toString
        |> text


findSecondSeen : List Int -> Int -> Set Int -> Int
findSecondSeen freqs current seenSoFar =
    if Set.member current seenSoFar then
        current
    else
        let
            seen =
                Set.insert current seenSoFar

            current_ =
                current + (List.head freqs |> Maybe.withDefault 0)
        in
        findSecondSeen (shift freqs) current_ seen


shift list =
    List.append (List.drop 1 list) (List.take 1 list)

-- Get rest of data here: https://adventofcode.com/2018/day/1/input
data =
    """
+13
-18
+13
+10
"""