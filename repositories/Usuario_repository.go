package repositories

import (
	"context"
	"time"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsuarioRepository struct {
    Collection *mongo.Collection
}

func NewUsuarioRepository(collection *mongo.Collection) *UsuarioRepository {
    return &UsuarioRepository{
        Collection: collection,
    }
}

func (r *UsuarioRepository) Create(usuario *models.Usuario) error {
    // Inicializar FechaRegistro automáticamente
    usuario.FechaRegistro = time.Now()

    _, err := r.Collection.InsertOne(context.Background(), usuario)
    return err
}

func (r *UsuarioRepository) FindAll() ([]*models.Usuario, error) {
    cursor, err := r.Collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var usuario []*models.Usuario
    if err = cursor.All(context.Background(), &usuario); err != nil {
        return nil, err
    }

    return usuario, nil
}

func (r *UsuarioRepository) FindByID(id bson.ObjectID) (*models.Usuario, error) {
    var usuario models.Usuario
    err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&usuario)
    if err != nil {
        return nil, err
    }
    return &usuario, nil
}

func (r *UsuarioRepository) Update(id bson.ObjectID, usuario *models.Usuario) error {
    _, err := r.Collection.UpdateOne(
        context.Background(),
        bson.M{"_id": id},
        bson.M{"$set": usuario},
    )
    return err
}

func (r *UsuarioRepository) Delete(id bson.ObjectID) error {
    _, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
    return err
}
