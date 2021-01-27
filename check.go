ALGORITHME Checkpoint 

var i,j,v: int;
var etat: bool;
var t: string;

BEGIN
read(t);

i:=0;
j:=0;
v:=0;
etat=true;

while(etat=true)

	while(t(i)<>" " and t[i]<>".") 					/*boucle de parcours du tab*/

		if (t[i] ="a" || t[i] ="e" || t[i] ="i" || 
			 t[i] ="u" || t[i] ="o" || t[i] ="y"){	/*test est ce qu il est voyelle ou non*/
			v=v+1
		}

		i=i+1;
	end while

	if (t[i]=' ') {				/*test lors de la sortie de la boucle a cause de lespace ou le point*/							*/
		j:=j+1;					/*j incremente le compteur des mots et je poursuis la boucle de parcours */	
		i=i+1;
	} else{					/*sinon la phrase est terminee je sors de la boucle de parcours et jaffiche le resultat*/	
		etat=false;
	}

end while

write("la longueur de la phrase est:" & i+1);
write ("dans cette phrase il y a" & j+1 & "mots");
write("il y a" & v & "voyelles");

END
