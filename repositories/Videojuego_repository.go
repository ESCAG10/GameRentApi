package repositories

import (
	"context"
	"time"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type VideojuegoRepository struct {
	Collection *mongo.Collection
}

func NewVideojuegoRepository(collection *mongo.Collection) *VideojuegoRepository {
	return &VideojuegoRepository{
		Collection: collection,
	}
}


// Crear videojuego
func (r *VideojuegoRepository) Create(videojuego *models.Videojuego) error {
	
	ahora := time.Now().Format("2006-01-02 15:04:05")
	
	videojuego.FechaCreacion = ahora
	videojuego.FechaActualizacion = ahora
	
	
	_, err := r.Collection.InsertOne(
		context.Background(),
		videojuego,
	)
	
	
	return err
}



// Obtener todos los videojuegos
func (r *VideojuegoRepository) FindAll() ([]models.Videojuego, error) {
	
	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)
	
	if err != nil {
		return nil, err
	}
	
	
	var documentos []bson.M
	if err := cursor.All(
		context.Background(),
		&documentos,
		); err != nil {

			return nil, err
		}
		
		var videojuegos []models.Videojuego
		
		for _, doc := range documentos {

			
			videojuego := models.Videojuego{}
			
			if id, ok := doc["_id"].(bson.ObjectID); ok {
				
				videojuego.ID = id
			
			}
			
			videojuego.Titulo, _ = doc["titulo"].(string)
			videojuego.Plataforma, _ = doc["plataforma"].(string)
			videojuego.Descripcion, _ = doc["descripcion"].(string)

			
			switch v := doc["precioRenta"].(type) {
				
			case float64:
				videojuego.PrecioRenta = v
			
			case int32:
				videojuego.PrecioRenta = float64(v)
			
			case int64:
				videojuego.PrecioRenta = float64(v)
			
			}
			
			switch v := doc["stock"].(type) {
				
			case int32:
				videojuego.Stock = int(v)
			
			case int64:
				videojuego.Stock = int(v)
			
			case int:
				videojuego.Stock = v
			
			}
			
			
			if activo, ok := doc["activo"].(bool); ok {
				videojuego.Activo = activo
			}
			
			if categoria, ok := doc["categoriaId"].(string); ok {
				
				videojuego.CategoriaID = categoria
			
			}

			
			videojuego.ImagenPortada, _ = doc["imagenPortada"].(string)
			videojuego.Desarrollador, _ = doc["desarrollador"].(string)
			videojuego.Editor, _ = doc["editor"].(string)
			videojuego.FechaCreacion, _ = doc["fechaCreacion"].(string)
			videojuego.FechaActualizacion, _ = doc["fechaActualizacion"].(string)
			
			
			
			videojuegos = append(videojuegos, videojuego)
		}
		
		
	return videojuegos, nil
}




// Obtener videojuego por ID
func (r *VideojuegoRepository) FindByID(id string) (*models.Videojuego, error) {
	
	objectID, err := bson.ObjectIDFromHex(id)
	
	
	if err != nil {
		
		return nil, err
	
	}
	
	
	
	var doc bson.M

	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
		).Decode(&doc)

		
		if err != nil {
			
			return nil, err
		
		}
		
		
		videojuego := models.Videojuego{}
		
		if id, ok := doc["_id"].(bson.ObjectID); ok {
			
			videojuego.ID = id
		
		}
		
		videojuego.Titulo, _ = doc["titulo"].(string)
		videojuego.Plataforma, _ = doc["plataforma"].(string)
		videojuego.Descripcion, _ = doc["descripcion"].(string)
		
		
		switch v := doc["precioRenta"].(type) {

			case float64:
				videojuego.PrecioRenta = v
				
			case int32:
				videojuego.PrecioRenta = float64(v)
				
			case int64:
				videojuego.PrecioRenta = float64(v)
			
			}


			
			switch v := doc["stock"].(type) {
				
			case int32:
				videojuego.Stock = int(v)
			case int64:
				videojuego.Stock = int(v)
			case int:
				videojuego.Stock = v
				
			}
			
			if activo, ok := doc["activo"].(bool); ok {
				
				videojuego.Activo = activo
			
			}
			
			if categoria, ok := doc["categoriaId"].(string); ok {
				
				videojuego.CategoriaID = categoria
			
			}
			
			videojuego.ImagenPortada, _ = doc["imagenPortada"].(string)
			videojuego.Desarrollador, _ = doc["desarrollador"].(string)
			videojuego.Editor, _ = doc["editor"].(string)
			videojuego.FechaCreacion, _ = doc["fechaCreacion"].(string)
			videojuego.FechaActualizacion, _ = doc["fechaActualizacion"].(string)
			
		return &videojuego, nil
		
}





// Actualizar videojuego
func (r *VideojuegoRepository) Update(id string, videojuego *models.Videojuego) error {
	
	objectID, err := bson.ObjectIDFromHex(id)
	
	if err != nil {
		return err

	}
	
	existente, err := r.FindByID(id)
	
	if err != nil {
		
		return err
	
	}
	
	videojuego.FechaCreacion = existente.FechaCreacion
	videojuego.FechaActualizacion = time.Now().Format("2006-01-02 15:04:05")
	
	
	_, err = r.Collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
		bson.M{
			"$set": bson.M{

				"titulo":             videojuego.Titulo,
				"plataforma":         videojuego.Plataforma,
				"descripcion":        videojuego.Descripcion,
				"precioRenta":        videojuego.PrecioRenta,
				"stock":              videojuego.Stock,
				"activo":             videojuego.Activo,
				"categoriaId":        videojuego.CategoriaID,
				"imagenPortada":      videojuego.ImagenPortada,
				"desarrollador":      videojuego.Desarrollador,
				"editor":             videojuego.Editor,
				"fechaCreacion":      videojuego.FechaCreacion,
				"fechaActualizacion": videojuego.FechaActualizacion,
			
			},
		},
	)
	
	return err

}

// Eliminar videojuego
func (r *VideojuegoRepository) Delete(id string) error {
	
	objectID, err := bson.ObjectIDFromHex(id)
	
	if err != nil {
		
		return err
	}
	
	_, err = r.Collection.DeleteOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	)
	
	return err

}
