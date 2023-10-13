sumlist([], 0).   % Caso base: la suma de una lista vac�a es 0.
sumlist([H|T], S) :- sumlist(T, Rest), S is H + Rest.
