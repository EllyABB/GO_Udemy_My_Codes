# Go_Project
Descarga y cursos del lenguaje de programación GO

## Instalación:
Link: 
https://go.dev/doc/install
Para linux:
Eliminar cualquier versión previa:
> $ rm -rf /usr/local/go 

Copiar de la carperta de Downloads a la requerida
> $ sudo cp Downloads/file /usr/local/go

Descomprimir (primero acceder como root)
Para acceder como usuario root tuve que desbloquear la cuenta root
> $ sudo passwd root
Y ahora si me sirve el comando:
> $ su -
para acceder como root

Entonces me muevo a la carpeta y descomprimo
> $ tar xzf go1.20.3.linux-amd64.tar.gz

Para más info (https://how-to.fandom.com/wiki/How_to_untar_a_tar_file_or_gzip-bz2_tar_file)

Para salir del usuario root:
> $ exit


Necesito hacer esto siempre: 
> export PATH=$PATH:/usr/local/go/bin  

Para evitarlo, tuve que editar el archivo:   
> cd /etc/  
> nano bash.bashrc  
Poner en la primera linea el anterior export. y guardarlo.



## Compilación
> $ go run name.go  
También, para construir un ejecutable  
>$ go build name.go


# Jupyter in go!
Se necesita instalar. Para linux funciona perfectamente:
>   go install github.com/gopherdata/gophernotes@v0.7.5  
  mkdir -p ~/.local/share/jupyter/kernels/gophernotes  
  cd ~/.local/share/jupyter/kernels/gophernotes  
  cp "$(go env GOPATH)"/pkg/mod/github.com/gopherdata/gophernotes@v0.7.5/kernel/*  "."  
  chmod +w ./kernel.json # in case copied kernel.json has no write permission  
  sed "s|gophernotes|$(go env GOPATH)/bin/gophernotes|" < kernel.json.in > kernel.json  
    
    
    
    
      
Pero para más info, con videos muy bonitos: https://github.com/gopherdata/gophernotes
