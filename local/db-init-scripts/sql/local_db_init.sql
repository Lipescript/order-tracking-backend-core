db("rgo001_users_active_delivery")
.createCollection('user_deliveries', {
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

