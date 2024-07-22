# Template integración continua AWS SAM y Golang

Este es un template para que futuros desarrolladores puedan descargar el código e implementar los microservicios necesarios

```bash
.
├── Makefile                    <-- Make para automatizar el build del codigo go (Es como el de c++)
├── README.md                   <-- Archivo de información
├── api                         <-- Código que se ejecutará en la lambda
│   ├── main.go                 <-- Código source de la lamda
│   └── main_test.go            <-- Test unitarios del main (si se hacen cambios en el input/output de la funcion, hay que cambiarlo, así mismo, aquí se deben probar las funciones)

└── template.yaml
```

## Requeridos

* AWS CLI configurado con permisos de administración
* [Docker instalado](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Trabajar en la lamda
  
* Clonar repositorio
  ```bash
  git clone https://github.com/GersonAM-Medismart/CICD-SAM-AWS-Template.git
  ```

* Crear un freatured para trabajar en el
  ```bash
  git checkout -b feature/nombre-de-la-rama

  ```
*Hacer los cambios necesarios en ./api

* Probar localmente: asegurarse que los cambios pasen las pruebas de manera local, usando ```go test```

* Hacer commit 
  ```bash
  git add .
  git commit -m "Descripción de los cambios realizados"
  git push origin feature/nombre-de-la-rama

  ```
*Hacer pull request y unir a main

Con esto se podría trabajar en el repositorio y los cambios se subirían automáticamente.

