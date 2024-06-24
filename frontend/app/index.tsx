import React, { useEffect, useState } from 'react';
import { View, Text, Button, StyleSheet, ActivityIndicator } from 'react-native';
import { GetUserLocation } from '../components/getUserLocation';
import { PrintShop } from '../components/printShop';
import { Shop } from '../components/Shop';
import {styles} from "../components/appStyles";

const App = () => {
  const [shopData, setShopData] = useState<{ shops: Shop[] } | null>(null);
  const [loading, setLoading] = useState(false);

  const fetchShop = async () => {
    setLoading(true);
    try {
      const location = await GetUserLocation();
      console.log(location);
      const response = await fetch('http://localhost:8080/api/choiceShop', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(location),
      });
      const data = await response.json();
      setShopData(data);
      console.log('data:', data);
    } catch (error) {
      console.error('error fetch shop:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <View style={styles.indexContainer}>
      <Text style={styles.header}>今日のお昼はここにしよう</Text>
      <Button title="さがす" onPress={fetchShop} />
      <View style={styles.bannerContainer}>
        <a href="http://webservice.recruit.co.jp/">
            <img src="http://webservice.recruit.co.jp/banner/hotpepper-s.gif" alt="ホットペッパーグルメ Webサービス" width="135" height="17" title="ホットペッパーグルメ Webサービス" />
        </a>
        </View>
      <View style={styles.printShop}>
        {loading ? (
          <ActivityIndicator size="large" color="#0000ff" />
        ) : (
          shopData && <PrintShop shopData={shopData} />
        )}
      </View>
    </View>
  );
};

export default App;
