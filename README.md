# PadelSort
## Proyecto de prácticas CC
Sistema para la gestión de clubs de pádel y jugadores.

[Configuración de entorno](https://github.com/carlostorralba/padelSort/blob/main/doc/config/hito0.md)
### Problema
Hoy en día no se disponen de muchas aplicaciones o webs para la gestión de clubs pádel y jugadores de pádel. 
Un club de pádel no dispone de la suficiente acerca de los jugadores de pádel que habitualmente juegan en el club o posibles nuevos jugadores, esta información que quiere conocer el club de pádel corresponde al nombre completo, *el nivel o categoría* del jugador y la zona en la que juega el jugador. Los clubes de pádel quieren esta información para llevar un recuento socios. 
A su vez los jugadores de pádel quieren conocer las categorías que se juegan en el club de pádel como el lugar y el porcentaje de jugadores por cada categoría.

Las categorías y niveles en relación al mundo del pádel pueden variar de un club a otro, es decir, es posible que un club de Granada tenga cinco categorías, siendo la 1º categoría la de mayor nivel, y otro club de Madrid tenga 7º categorías, siendo la 7º la de mayor nivel

### Solución

Por ello se propone implementar un sistema que sea capaz de llevar la gestión de los clubs de pádel así como permita a los jugadores de pádel registrarse en clubs de pádel.
Para el registro de un club de pádel será necesario conocer los siguientes datos:
 * Lugar
 * Números de teléfonos
 * Categorías
 * Porcentaje de jugadores por cada categoría.

Esta información del club de pádel le será muy útil al jugador de pádel que quiera incribirse en el club ya que podrá decidir que club de pádel es mejor para sus carácteristicas a su vez también le será útil al jugador de pádel que ya es socio de un club de pádel para ver como avanza el porcentaje de jugadores por cada categoría y de esta manera si el sube de nivel o baja tendrá la opción de inscribirse en otro club el cual mejore las condiciones de este.

Para el registro de un jugador de pádel será necesario conocer los siguientes datos:
* Nombre completo
* Nombre de usuario (opcional)
* Número de teléfono (opcional)
* Nivel de pádel asociado a cada club.
* Zona en la que suele jugar frecuentemente.

Esta información será valiosa para el club de pádel ya que podrá llevar un registro de sus jugadores y de esta manera contabilizarlos y agruparlos por categorías. Todos estos datos son importantes pero sobre los datos del nivel de pádel ya que permite clasificar al jugador así como la zona en la que juega para que el club sepa si es un jugador que estará de paso o es posible que sea un cliente fijo. El jugador de pádel podrá inscribirse en cuantos club de pádel quiera y a su vez tendrá asociado un nivel de pádel por cada club debido a que el nivel en cada club varía. En el caso de que un jugador se registre en el sistema deberá introducir obligatorimente los datos de su nombre completo y la zona en la que suele jugar y por consiguiente si se quiere registrar en un club de pádel se le solicitará una prueba de nivel con el club de pádel para determinar el nivel del jugador en ese club.

