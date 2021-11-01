# Lógica de negocio

El sistema requerirá información acerca de las características físicas de los usuarios para que contribuyan al sistema de recomendación de rutinas. Este sistema de recomendación 
será implementado utilizando un algoritmo de machine learning. El sistema de recomendación utilizará las características físicas de los usuarios para recomendar rutinas a otros 
usuarios con características similares.

Otro aspecto que tendrá en cuenta el sistema de recomendación es la rutina que quiere el usuario, es decir, el usuario podrá filtrar el tipo de rutina (Ej: para ganar masa muscular
, nivel de experiencia, días de entrenamiento) que quiere para de esta manera afinar más la búsqueda del sistema de recomendación. 
También cabe destacar que el sistema de recomendación tendrá en cuenta las rutinas más seguidas por los usuarios.
Las características físicas de los usuarios serán validadas conforme a unos estándares, ya que un usuario no podrá introducir como característica suya una altura de 1000 m. 
Por consiguiente, la información de las rutinas también será validada, es decir, no podrán contener entrenamientos de más de 6 días y más de 10 ejercicios por entrenamiento 
para que de esta manera sea una rutina óptima y eficiente. A su vez los ejercicios no podrán contener más de 10 series ni más de 50 repeticiones, 
por otra parte el descanso por repetición deberá ser mínimamente de 30 segundos y no más de 5 minutos (300 segundos), todo ello para que se cumpla lo descrito anteriormente.
