import React, { useState, useEffect } from "react";
import {View,Text,FlatList,ActivityIndicator,StyleSheet} from "react-native";
import axios from "axios";

export default function App() {
      const [usuarios, setUsuarios] = useState([]);
      const [loading, setLoading] = useState(true); // Estado de carregamento
      const [error, setError] = useState(null);
      useEffect(() => {
        const fetchUsuarios = async () => {
          try {
            const response = await axios.get(
              "https://jsonplaceholder.typicode.com/users"
            );
            setUsuarios(response.data);
          } catch (error) {
            console.error("Erro ao carregar usuários:", error);
            setError("Erro ao carregar usuários.");
          } finally {
            setLoading(false); // Atualiza o estado de carregamento para false
          }
        };
        fetchUsuarios();
      }, []);
      if (loading) {
        return (
          <View style={styles.container}>
            <ActivityIndicator size="large" color="#0000ff" />
            <Text style={styles.loading}>Carregando...</Text>
          </View>
        );
      }
      if (error) {
        return (
          <View style={styles.container}>
            <Text style={styles.error}>{error}</Text>
          </View>
        );
      }
      return (
        <View style={styles.container}>
          <Text style={styles.title}>Lista de Usuários</Text>
          <FlatList
            data={usuarios}
            keyExtractor={(item) => item.id.toString()} // Gera uma chave única 
            renderItem={({ item }) => (
              <Text style={styles.item}>{item.name}</Text>
            )}
          />
        </View>
      );
}
const styles = StyleSheet.create({
      container: {
        flex: 1,
        padding: 16,
        backgroundColor: "#f5f5f5",
      },
      loading: {
        marginTop: 10,
        fontSize: 18,
        color: "#555",
      },
      error: {
        color: "red",
        fontSize: 16,
        textAlign: "center",
      },
      title: {
        fontSize: 24,
        fontWeight: "bold",
        marginBottom: 10,
        textAlign: "center",
      },
      item: {
        fontSize: 18,
        padding: 10,
        borderBottomWidth: 1,
        borderBottomColor: "#ccc",
      },
});