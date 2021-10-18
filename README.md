# RouGym
## Proyecto de prácticas CC
RouGym (Routine Gym) será un sistema de recomendación de rutinas de entrenamiento. 

[Configuración de entorno](https://github.com/carlostorralba/padelSort/blob/main/doc/config/hito0.md)

[Documentación adicional hito 1](https://github.com/carlostorralba/RouGym/blob/hito1/doc/hito1.md)

### Problema
En la actualidad muchas personas que quieren ir o van al gimnasio suelen tener problemas a la hora de encontrar una rutina de entrenamiento óptima acorde con sus características. Muchos de los gimnasios suelen disponer de rutinas de entrenamiento genéricas pero estas rutinas no se adaptan a las características de cada persona. Por esto muchas personas se ven obligadas a contratar a un entrenador personal para que les realice una rutina de entrenamiento personalizada.

Otro problema es que los usuarios no tienen un historial de las rutinas utilizadas a lo largo del tiempo, y no pueden visualizar su progreso de rutinas.

### Solución
Por ello se propone un sistema de recomendación de rutinas de entrenamiento en el cual a partir de las características del usuario, como el tipo de entrenamiento que quiera, experiencia en gimnasios y sus atributos físicos, se le recomendarán un tipo de rutinas u otras.  Los usuarios podrán publicar sus rutinas y que estas a su vez podrán ser seguidas por otros usuarios. El sistema incluirá un sistema de cobro para acceder a determinadas rutinas, es decir, las rutinas con un elevado número de seguidores tendrán un importe para acceder a ellas. Los usuarios los cuales contengan rutinas con un elevado número de seguidores recibirán un beneficio económico. 

Con respecto a lo anterior los usuarios únicamente podrán crear como máximo dos rutinas, en caso de que el usuario desee añadir más rutinas deberá efectuar un pago por ello.

El sistema le proporcionará también al usuario un historial de las rutinas utilizadas, es decir, se le mostrará a través de una gráfica del tipo de rutinas utilizadas a lo largo del tiempo para que pueda visualizar su progreso.


### Beneficio
Este sistema será beneficioso para todos los tipos de personas que van al gimnasio ya que podrán encontrar una rutina acorde con sus características y necesidades sin tener que perder mucho tiempo buscando por internet y sin tener que contratar a un entrenador personal para la realización de una rutina personalizada. A su vez los usuarios los cuales tengan rutinas con un elevado número de seguidores recibirán una beneficio económico por ello.

### Lógica de negocio

El sistema requerirá información acerca de las características físicas de los usuarios para que contribuyan al sistema de recomendación de rutinas. Este sistema de recomendación será implementado utilizando un algoritmo de machine learning. El sistema de recomendación utilizará las características físicas de los usuarios para recomendar rutinas a otros usuarios con características similares.

Otro aspecto que tendrá en cuenta el sistema de recomendación es la rutina que quiere el usuario, es decir, el usuario podrá filtrar el tipo de rutina (Ej: para ganar masa muscular, nivel de experiencia, días de entrenamiento) que quiere para de esta manera afinar más la búsqueda del sistema de recomendación. A su vez el sistema de recomendación tendrá en cuenta el historial de rutinas realizas por el usuario para mejorar la búsqueda. También cabe destacar que el sistema de recomendación tendrá en cuenta las rutinas más seguidas por los usuarios.
Las características físicas de los usuarios serán validadas conforme a unos estándares, ya que un usuario no podrá introducir como característica suya una altura de 1000 m. Por consiguiente, la información de las rutinas también será validada, es decir, no podrán contener entrenamientos de más de 6 días y más de 10 ejercicios por entrenamiento.


### Nube
Este sistema será beneficioso para la nube ya que podrá haber múltiples usuarios requiriendo y compartiendo rutinas a cualquier hora del día los 365 días al año, con lo cual se necesitará que esa información este almacenada en la nube para que sea accesible de manera rápida y óptima para muchos usuarios. También será desplegado en la nube ya que podrá ser utilizado desde cualquier dispositivo y desde cualquier lugar.


### Lenguaje
Se va utilizar el lenguaje Go porque es un lenguaje compilado de código abierto, que está orientado a objetos y es mutiplataforma. He escogido este lenguaje ya que lo he utilizado en una asignatura de mi carrera y quiero ampliar conocimientos en él.


