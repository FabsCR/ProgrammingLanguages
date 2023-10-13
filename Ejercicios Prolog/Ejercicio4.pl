partir([], _, [], []).  % Caso base: la lista vacía resulta en dos listas vacías.

partir([X|Xs], Umbral, [X|Menores], Mayores) :-
    X =< Umbral,
    partir(Xs, Umbral, Menores, Mayores).

partir([X|Xs], Umbral, Menores, [X|Mayores]) :-
    X > Umbral,
    partir(Xs, Umbral, Menores, Mayores).
