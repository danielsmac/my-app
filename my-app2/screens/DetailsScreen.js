import React from "react";
import { View, Text, Button } from "react-native";
function DetailsScreen({ route, navigation }) {
    // Desestrutura os parâmetros recebidos da rota
    const { itemid, otherParam } = route.params;

      return (
        <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
          <Text>Details Screen</Text>
          <Text>ID: {itemid}</Text>
          <Text>Other Parameter: {otherParam}</Text>
          <Button title="Go to Home" onPress={() => navigation.navigate("Home")} />
        </View>
      );
}
export default DetailsScreen;