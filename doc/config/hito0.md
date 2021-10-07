# Configuración del entorno hito 0
## Creación de claves SSH y subida a GitHub
Para está configuración se han seguido los pasos de está página de Github Docs: https://docs.github.com/en/authentication/connecting-to-github-with-ssh
### Creación de claves SSH
Se crean una par de claves asociadas a la cuenta de Github.

![Alt Text](https://github.com/carlostorralba/padelSort/blob/hito0/doc/images/sshGenerate.PNG)

A contiuación se añade una identidad a las claves SSH que en esta caso es el correo de la cuenta Github.

![Alt Text](https://github.com/carlostorralba/padelSort/blob/hito0/doc/images/sshAddSshAgent.PNG)

### Subida a GitHub

A raíz de lo anterior se añade la clave pública a la cuenta al perfil de la cuenta Github.

![Alt Text](https://github.com/carlostorralba/padelSort/blob/hito0/doc/images/addSshKeyToGithub.PNG)

Y por último se comprueba que la conexión a Github a través de SSH funciona.

![Alt Text](https://github.com/carlostorralba/padelSort/blob/hito0/doc/images/TestConnectionSsh.PNG)

## Creación del respositorio y configuración del nombre, correo electrónico
### Creación del repositorio
Se crea el repositorio añadiendo los ficheros visualizados en la imagen:
  * README: que contendrá la explicación del proyecto.
  * .gitignore: el cual se ha seleccionado el lenguaje de implementación del proyecto.
  * license: fichero que contendrá la licencia seleccionada en la imagen.


 ![Alt Text](https://github.com/carlostorralba/padelSort/blob/hito0/doc/images/CreateRepository.PNG)


### Configuración del nombre y correo electrónico
Inicialmente se clona el repositorio creado y una vez hecho esto se configura el email y el usuario de manera local para ese repositorio.

![image](https://user-images.githubusercontent.com/91627117/136361824-e009c39f-6bda-4504-ab37-b786ff9bd2c7.png)

## Edición del perfil de Github para que aparezca nombre completo, avatar, ciudad y universidad

![Alt Text](https://github.com/carlostorralba/padelSort/blob/hito0/doc/images/UpdateProfileGithub.PNG)


## Segundo factor de autentificación
Para esta configuración hay que desplazarse a *Settings* del usuario y una vez aquí seleccionar *Account security* y añadir un segundo factor de autentificación.
Para añadirlo tienes dos opciones:
 *  A través de la app móvil.
 *  A través de SMS (mi caso).

![Alt Text](https://github.com/carlostorralba/padelSort/blob/hito0/doc/images/twoFactorAuth.PNG)

![Alt Text](https://github.com/carlostorralba/padelSort/blob/hito0/doc/images/twoFactorAuthActivated.PNG)

