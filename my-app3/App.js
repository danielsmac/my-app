import react from "react";
import { View, Text, StyleSheet } from "react-native";
import axios from "axios";

function App () {
  async function fetchdata () {
    try {
      const response = await axios.get("https://jsonplaceholder.typicode.com/users");
      console.log(response.data); // Exibe os dados no console
    } catch (error) {
      console.error('Erro na requisição', error); // Exibe erros no console
    }
  }
  fetchdata();

  return (
    <View style={styles.container}>
      <Text style ={styles.title}>Exemplo de uso do Axios</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#F5F5F5',
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
  },
});

export default App;