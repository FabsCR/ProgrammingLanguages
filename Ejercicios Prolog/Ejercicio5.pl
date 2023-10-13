% Define el predicado sub_string/3 para verificar si una cadena contiene una subcadena.
sub_string(Subcadena, Cadena) :- sub_string(Cadena, _, _, _, Subcadena).

% Define el predicado sub_cadenas/3 para filtrar las cadenas que contienen una subcadena.
sub_cadenas(_, [], []).  % Caso base: la lista vacía resulta en una lista vacía.

sub_cadenas(Subcadena, [Cadena|Resto], Filtradas) :-
    (   sub_string(Subcadena, Cadena)
    ->  Filtradas = [Cadena|RestoFiltradas]
    ;   Filtradas = RestoFiltradas
    ),
    sub_cadenas(Subcadena, Resto, RestoFiltradas).
