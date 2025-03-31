require('dotenv').config();
const express = require('express');
const { MongoClient } = require('mongodb');

const app = express();
const port = process.env.PORT || 3035;
const uri = process.env.MONGODB_URI || 'mongodb://root:1234@localhost:27017/?authSource=admin';
const client = new MongoClient(uri); 

// Função para inicializar o banco de dados e as coleções
async function initializeDatabase() {
  try {
    // Conecta ao MongoDB
    await client.connect();
    console.log('Conectado ao MongoDB');

    // Seleciona ou cria o banco de dados
    const db = client.db('rgo001_users_active_delivery');

    // Cria a collection 'user_deliveries' com validação de schema
    const collection = await db.createCollection('user_deliveries', {
      validator: {
        $jsonSchema: {
          bsonType: 'object',
          required: ['user_id', 'delivery_id', 'delivery_status', 'message_timestamp'],
          properties: {
            user_id: { bsonType: 'string' },
            delivery_id: { bsonType: 'string' },
            delivery_status: {
              enum: ['pendente', 'em_transito', 'entregue', 'cancelado']
            },
            message_timestamp: { bsonType: 'date' }
          }
        }
      }
    });

    // Cria o índice na coleção
    await collection.createIndex({ user_id: 1 }, { unique: true });
    await collection.createIndex({ delivery_id: 1 }, { unique: true });
    await collection.createIndex({ user_id: 1, delivery_id: 1 }, { unique: true });

    console.log('Collection "user_deliveries" criada com sucesso.');
  } catch (error) {
    console.error('Erro ao criar banco de dados e collection:', error);
  } finally {
    // Fecha a conexão com o MongoDB
    if (client) {
      await client.close();
      console.log('Conexão com o MongoDB fechada.');
    }
  }
}
// Rota para inicializar o banco de dados e as coleções
app.get('/create-database', async (req, res) => {
  await initializeDatabase();
  res.status(200).send('Banco de dados e collection criados com sucesso.');
});


// Inicia o servidor
app.listen(port, () => {
  console.log(`Servidor rodando em http://localhost:${port}`);
});