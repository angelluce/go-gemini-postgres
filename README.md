# Go Gemini PostgreSQL API

Una API RESTful desarrollada en Go que integra el modelo de lenguaje Gemini de Google con una base de datos PostgreSQL. Esta API permite gestionar ítems y generar descripciones de texto o audio (voz) para ellos, utilizando la inteligencia artificial de Gemini y la síntesis de voz de Google Cloud Text-to-Speech.

## Características

* **Gestión de ítems:** Almacena y recupera información de ítems desde una base de datos PostgreSQL.
* **Integración con Gemini:** Genera respuestas de texto coherentes y contextualmente relevantes utilizando el modelo de lenguaje Google Gemini.
* **Síntesis de voz (Text-to-Speech):** Convierte las respuestas de texto de Gemini en archivos de audio (MP3) utilizando la API de Google Cloud Text-to-Speech.
* **API RESTful:** Endpoints bien definidos para una fácil interacción.
* **Configuración por variables de entorno:** Gestión segura de credenciales y configuraciones.

## Requisitos

Antes de ejecutar este proyecto, asegúrate de tener instalado lo siguiente:

* **Go:** Versión 1.23.4 o superior.
* **PostgreSQL:** Base de datos activa y accesible.
* **Cuenta de Google Cloud:** Con un proyecto y las APIs necesarias habilitadas.
    * **Generative Language API** (para Gemini).
    * **Cloud Text-to-Speech API**.
* **Claves de API de Google Cloud:** Con permisos para ambas APIs mencionadas.

## Configuración del Proyecto

### 1. Clonar el repositorio

```bash
git clone https://github.com/angelluce/go-gemini-postgres.git
cd go-gemini-postgres-api
```

### 2. Configuración de variables de entorno (.env)

Crea un archivo llamado .env en la raíz de tu proyecto con las siguientes variables:

```text
DATABASE_URL="postgres://user:password@host:port/database?sslmode=disable"
GEMINI_API_KEY="TU_CLAVE_API_DE_GEMINI_AQUI"
GOOGLE_CLOUD_TTS_API_KEY="TU_CLAVE_API_DE_GOOGLE_CLOUD_TTS_AQUI"
PORT=":8080"
```

- DATABASE_URL: La cadena de conexión a tu base de datos PostgreSQL. Reemplazar user, password, host, port y database con tus credenciales reales. sslmode=disable para desarrollo local.

- GEMINI_API_KEY: Tu clave de API para la Google Generative Language API. Puedes generar una desde Google AI Studio o la consola de Google Cloud.

- GOOGLE_CLOUD_TTS_API_KEY: Tu clave de API para la Google Cloud Text-to-Speech API. Asegúrate de que esta clave tenga permisos para la API de TTS y de que la API esté habilitada en tu proyecto de Google Cloud.

- PORT: El puerto en el que la API escuchará las solicitudes.

### 3. Configuración de la Base de Datos (PostgreSQL)

Asegúrate de tener una tabla en tu base de datos PostgreSQL,puedes usar esta estructura:

```sql
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

INSERT INTO items (name, description) VALUES
('Laptop Gaming', 'Una potente laptop diseñada para juegos con gráficos RTX 4080 y 32GB de RAM DDR5.'),
('Teclado Mecánico RGB', 'Teclado con switches Cherry MX Brown, retroiluminación RGB personalizable y diseño ergonómico.'),
('Mouse Inalámbrico Ergonómico', 'Mouse de alta precisión con sensor óptico de 16000 DPI y conectividad inalámbrica 2.4 GHz y Bluetooth.');
```

### 4. Instalar Dependencias

En la raíz del proyecto desde la terminal, descarga las dependencias de Go:

```bash
go mod tidy
```

## Ejecutar la Aplicación
Una vez configurado, puedes iniciar la API:

```bash
go run main.go
```

La API se iniciará y escuchará en el puerto especificado en tu archivo `.env` (por defecto, `localhost:8080`).

## Ejecutar la Aplicación con Docker

El proyecto implementó la técnica multi-stage build con la finalidad optimizar el tamaño de la imágen final

```bash
docker compose up -d

docker logs go-gemini-postgres-app-1

docker compose down
```

## Endpoints de la API

La API expone los siguientes endpoints:

1. **GET /api/items**: Obtiene una lista de todos los ítems almacenados en la base de datos.

- Método: GET
- URL: http://localhost:8080/api/items
- Respuesta exitosa (200 OK):

```json
[
  {
    "id": 1,
    "name": "Laptop Gaming",
    "description": "Una potente laptop diseñada para juegos con gráficos RTX 4080 y 32GB de RAM DDR5."
  },
  {
    "id": 2,
    "name": "Teclado Mecánico RGB",
    "description": "Teclado con switches Cherry MX Brown, retroiluminación RGB personalizable y diseño ergonómico."
  }
]
```

2. **POST /api/generate/text**: Genera una descripción de texto para un ítem específico utilizando Google Gemini. El prompt principal está en el código del servidor, solo necesitas proporcionar el item_id.

- Método: POST
- URL: http://localhost:8080/api/generate/text
- Cuerpo de la solicitud (JSON):

```json
{
  "item_id": 1
}
```

- Respuesta exitosa (200 OK):

```json
{
  "response": "Este portátil es una bestia de rendimiento con 16GB de RAM y 512GB de SSD, ideal para tareas exigentes y multitarea fluida. Su amplio espacio de almacenamiento te permitirá guardar todos tus archivos importantes con facilidad. Es la elección perfecta para quienes buscan velocidad y eficiencia."
}
```

3. **POST /api/generate/audio**: Genera una descripción de audio (MP3) para un ítem específico utilizando Google Gemini y Google Cloud Text-to-Speech. El prompt principal está en el código del servidor, solo necesitas proporcionar el item_id.

- Método: POST
- URL: http://localhost:8080/api/generate/audio
- Cuerpo de la solicitud (JSON):

```json
{
  "item_id": 2
}
```

- Respuesta exitosa (200 OK): 

La API devolverá un stream de bytes MP3. 

Para guardar el archivo de audio con curl:

```bash
curl -X POST \
  http://localhost:8080/api/generate/audio \
  -H 'Content-Type: application/json' \
  -d '{"item_id": 2}' \
  --output item_description.mp3
```

Para guardar el archivo de audio con Postman: Después de enviar la solicitud, busca la opción "Save response" o el icono de descarga en la sección de respuesta y guárdalo como .mp3.

## Estructura del Proyecto

```text
go-gemini-postgres-api/
├── main.go                       # Punto de entrada de la aplicación, configuración de rutas.
├── config/
│   └── config.go                 # Carga de variables de entorno.
├── database/
│   └── postgres.go               # Conexión y operaciones con la base de datos PostgreSQL.
├── handlers/
│   └── handlers.go               # Lógica de los endpoints de la API.
├── models/
│   └── item.go                   # Definición de la estructura del modelo Item.
├── services/
│   ├── gemini.go                 # Integración con la API de Google Gemini.
│   └── tts_google_cloud.go       # Integración con la API de Google Cloud Text-to-Speech.
├── .gitignore                    # Archivos y directorios a ignorar por Git.
├── go.mod                        # Módulos de Go del proyecto.
└── go.sum                        # Sumas de verificación de los módulos.
```

## Contribuciones

Si deseas contribuir a este proyecto, no dudes en abrir un *issue* o enviar un *pull request*.
