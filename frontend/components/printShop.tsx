import React from 'react';
import { View, Text, StyleSheet, TouchableOpacity, Linking } from 'react-native';
import { Shop } from "./Shop"; 
import {styles} from "./appStyles";

interface PrintShopProps {
  shopData: { shops: Shop[] };
}

export const PrintShop: React.FC<PrintShopProps> = ({ shopData }) => {
    const openMap = (name: string) => {
      const url = `http://maps.google.com/?q=${encodeURIComponent(name)}`;
      Linking.openURL(url);
    };
  
    return (
      <View style={styles.printShopContainer}>
        {shopData.shops.map((shop, index) => (
          <View key={index} style={styles.card}>
            <Text style={styles.title}>{shop.Name}</Text>
            <TouchableOpacity onPress={() => openMap(shop.Name)}>
              <Text style={styles.url}>open with google map</Text>
            </TouchableOpacity>
            
            <Text style={styles.text}>{shop.Genre}</Text>
            <Text style={styles.text}>{shop.OpenHour}</Text>
          </View>
        ))}
      </View>
    );
  };
