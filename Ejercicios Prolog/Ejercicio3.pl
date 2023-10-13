aplanar([], []).  % Caso base: la lista vacía se aplanará como una lista vacía.
aplanar([X|Xs], L2) :- 
    aplanar(X, XFlat),     % Aplanar el primer elemento de la lista.
    aplanar(Xs, XsFlat),   % Aplanar el resto de la lista.
    append(XFlat, XsFlat, L2).  % Concatenar las listas aplanadas.
