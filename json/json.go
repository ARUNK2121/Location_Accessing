package json

// // You can edit this code!
// // Click here and start typing.
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {

// 	bigData := `{
//   "type": "FeatureCollection",
//   "query": [
//     -73.989,
//     40.733
//   ],
//   "features": [
//     {
//       "id": "address.7036282930256310",
//       "type": "Feature",
//       "place_type": [
//         "address"
//       ],
//       "relevance": 1,
//       "properties": {
//         "accuracy": "rooftop",
//         "mapbox_id": "dXJuOm1ieGFkcjo4NTA4ODFhOS1lZjY1LTRkMTUtOWQyYS1kNmY0MTllZmJlMDI"
//       },
//       "text": "East 13th Street",
//       "place_name": "120 East 13th Street, New York, New York 10003, United States",
//       "center": [
//         -73.98893,
//         40.73295
//       ],
//       "geometry": {
//         "type": "Point",
//         "coordinates": [
//           -73.98893,
//           40.73295
//         ]
//       },
//       "address": "120",
//       "context": [
//         {
//           "id": "neighborhood.187075820",
//           "mapbox_id": "dXJuOm1ieHBsYzpDeWFNN0E",
//           "wikidata": "Q1043326",
//           "text": "East Village"
//         },
//         {
//           "id": "postcode.24620780",
//           "mapbox_id": "dXJuOm1ieHBsYzpBWGV1N0E",
//           "text": "10003"
//         },
//         {
//           "id": "locality.338856684",
//           "mapbox_id": "dXJuOm1ieHBsYzpGREtLN0E",
//           "wikidata": "Q11299",
//           "text": "Manhattan"
//         },
//         {
//           "id": "place.233720044",
//           "mapbox_id": "dXJuOm1ieHBsYzpEZTVJN0E",
//           "wikidata": "Q60",
//           "text": "New York"
//         },
//         {
//           "id": "district.17000172",
//           "mapbox_id": "dXJuOm1ieHBsYzpBUU5tN0E",
//           "wikidata": "Q500416",
//           "text": "New York County"
//         },
//         {
//           "id": "region.107756",
//           "mapbox_id": "dXJuOm1ieHBsYzpBYVRz",
//           "wikidata": "Q1384",
//           "short_code": "US-NY",
//           "text": "New York"
//         },
//         {
//           "id": "country.8940",
//           "mapbox_id": "dXJuOm1ieHBsYzpJdXc",
//           "wikidata": "Q30",
//           "short_code": "us",
//           "text": "United States"
//         }
//       ]
//     },
//     {
//       "id": "neighborhood.187075820",
//       "type": "Feature",
//       "place_type": [
//         "neighborhood"
//       ],
//       "relevance": 1,
//       "properties": {
//         "mapbox_id": "dXJuOm1ieHBsYzpDeWFNN0E",
//         "wikidata": "Q1043326"
//       },
//       "text": "East Village",
//       "place_name": "East Village, New York, New York, United States",
//       "bbox": [
//         -73.991921,
//         40.725203,
//         -73.982557,
//         40.734816
//       ],
//       "center": [
//         -73.987361,
//         40.729269
//       ],
//       "geometry": {
//         "type": "Point",
//         "coordinates": [
//           -73.987361,
//           40.729269
//         ]
//       },
//       "context": [
//         {
//           "id": "postcode.24620780",
//           "mapbox_id": "dXJuOm1ieHBsYzpBWGV1N0E",
//           "text": "10003"
//         },
//         {
//           "id": "locality.338856684",
//           "mapbox_id": "dXJuOm1ieHBsYzpGREtLN0E",
//           "wikidata": "Q11299",
//           "text": "Manhattan"
//         },
//         {
//           "id": "place.233720044",
//           "mapbox_id": "dXJuOm1ieHBsYzpEZTVJN0E",
//           "wikidata": "Q60",
//           "text": "New York"
//         },
//         {
//           "id": "district.17000172",
//           "mapbox_id": "dXJuOm1ieHBsYzpBUU5tN0E",
//           "wikidata": "Q500416",
//           "text": "New York County"
//         },
//         {
//           "id": "region.107756",
//           "mapbox_id": "dXJuOm1ieHBsYzpBYVRz",
//           "wikidata": "Q1384",
//           "short_code": "US-NY",
//           "text": "New York"
//         },
//         {
//           "id": "country.8940",
//           "mapbox_id": "dXJuOm1ieHBsYzpJdXc",
//           "wikidata": "Q30",
//           "short_code": "us",
//           "text": "United States"
//         }
//       ]
//     },
//     {
//       "id": "postcode.24620780",
//       "type": "Feature",
//       "place_type": [
//         "postcode"
//       ],
//       "relevance": 1,
//       "properties": {
//         "mapbox_id": "dXJuOm1ieHBsYzpBWGV1N0E"
//       },
//       "text": "10003",
//       "place_name": "New York, New York 10003, United States",
//       "bbox": [
//         -73.9996048,
//         40.722932,
//         -73.9798632,
//         40.7396739
//       ],
//       "center": [
//         -73.989876,
//         40.732075
//       ],
//       "geometry": {
//         "type": "Point",
//         "coordinates": [
//           -73.989876,
//           40.732075
//         ]
//       },
//       "context": [
//         {
//           "id": "locality.338856684",
//           "mapbox_id": "dXJuOm1ieHBsYzpGREtLN0E",
//           "wikidata": "Q11299",
//           "text": "Manhattan"
//         },
//         {
//           "id": "place.233720044",
//           "mapbox_id": "dXJuOm1ieHBsYzpEZTVJN0E",
//           "wikidata": "Q60",
//           "text": "New York"
//         },
//         {
//           "id": "district.17000172",
//           "mapbox_id": "dXJuOm1ieHBsYzpBUU5tN0E",
//           "wikidata": "Q500416",
//           "text": "New York County"
//         },
//         {
//           "id": "region.107756",
//           "mapbox_id": "dXJuOm1ieHBsYzpBYVRz",
//           "wikidata": "Q1384",
//           "short_code": "US-NY",
//           "text": "New York"
//         },
//         {
//           "id": "country.8940",
//           "mapbox_id": "dXJuOm1ieHBsYzpJdXc",
//           "wikidata": "Q30",
//           "short_code": "us",
//           "text": "United States"
//         }
//       ]
//     },
//     {
//       "id": "locality.338856684",
//       "type": "Feature",
//       "place_type": [
//         "locality"
//       ],
//       "relevance": 1,
//       "properties": {
//         "mapbox_id": "dXJuOm1ieHBsYzpGREtLN0E",
//         "wikidata": "Q11299"
//       },
//       "text": "Manhattan",
//       "place_name": "Manhattan, New York, United States",
//       "bbox": [
//         -74.0473132,
//         40.679573,
//         -73.907,
//         40.882075
//       ],
//       "center": [
//         -73.959894,
//         40.789624
//       ],
//       "geometry": {
//         "type": "Point",
//         "coordinates": [
//           -73.959894,
//           40.789624
//         ]
//       },
//       "context": [
//         {
//           "id": "place.233720044",
//           "mapbox_id": "dXJuOm1ieHBsYzpEZTVJN0E",
//           "wikidata": "Q60",
//           "text": "New York"
//         },
//         {
//           "id": "district.17000172",
//           "mapbox_id": "dXJuOm1ieHBsYzpBUU5tN0E",
//           "wikidata": "Q500416",
//           "text": "New York County"
//         },
//         {
//           "id": "region.107756",
//           "mapbox_id": "dXJuOm1ieHBsYzpBYVRz",
//           "wikidata": "Q1384",
//           "short_code": "US-NY",
//           "text": "New York"
//         },
//         {
//           "id": "country.8940",
//           "mapbox_id": "dXJuOm1ieHBsYzpJdXc",
//           "wikidata": "Q30",
//           "short_code": "us",
//           "text": "United States"
//         }
//       ]
//     },
//     {
//       "id": "place.233720044",
//       "type": "Feature",
//       "place_type": [
//         "place"
//       ],
//       "relevance": 1,
//       "properties": {
//         "mapbox_id": "dXJuOm1ieHBsYzpEZTVJN0E",
//         "wikidata": "Q60"
//       },
//       "text": "New York",
//       "place_name": "New York, New York, United States",
//       "bbox": [
//         -74.2596399,
//         40.477399,
//         -73.700292,
//         40.917577
//       ],
//       "center": [
//         -73.9866,
//         40.7306
//       ],
//       "geometry": {
//         "type": "Point",
//         "coordinates": [
//           -73.9866,
//           40.7306
//         ]
//       },
//       "context": [
//         {
//           "id": "district.17000172",
//           "mapbox_id": "dXJuOm1ieHBsYzpBUU5tN0E",
//           "wikidata": "Q500416",
//           "text": "New York County"
//         },
//         {
//           "id": "region.107756",
//           "mapbox_id": "dXJuOm1ieHBsYzpBYVRz",
//           "wikidata": "Q1384",
//           "short_code": "US-NY",
//           "text": "New York"
//         },
//         {
//           "id": "country.8940",
//           "mapbox_id": "dXJuOm1ieHBsYzpJdXc",
//           "wikidata": "Q30",
//           "short_code": "us",
//           "text": "United States"
//         }
//       ]
//     },
//     {
//       "id": "district.17000172",
//       "type": "Feature",
//       "place_type": [
//         "district"
//       ],
//       "relevance": 1,
//       "properties": {
//         "mapbox_id": "dXJuOm1ieHBsYzpBUU5tN0E",
//         "wikidata": "Q500416"
//       },
//       "text": "New York County",
//       "place_name": "New York County, New York, United States",
//       "bbox": [
//         -74.0517904,
//         40.6692438,
//         -73.907007,
//         40.8853224
//       ],
//       "center": [
//         -73.965633,
//         40.777244
//       ],
//       "geometry": {
//         "type": "Point",
//         "coordinates": [
//           -73.965633,
//           40.777244
//         ]
//       },
//       "context": [
//         {
//           "id": "region.107756",
//           "mapbox_id": "dXJuOm1ieHBsYzpBYVRz",
//           "wikidata": "Q1384",
//           "short_code": "US-NY",
//           "text": "New York"
//         },
//         {
//           "id": "country.8940",
//           "mapbox_id": "dXJuOm1ieHBsYzpJdXc",
//           "wikidata": "Q30",
//           "short_code": "us",
//           "text": "United States"
//         }
//       ]
//     },
//     {
//       "id": "region.107756",
//       "type": "Feature",
//       "place_type": [
//         "region"
//       ],
//       "relevance": 1,
//       "properties": {
//         "mapbox_id": "dXJuOm1ieHBsYzpBYVRz",
//         "wikidata": "Q1384",
//         "short_code": "US-NY"
//       },
//       "text": "New York",
//       "place_name": "New York, United States",
//       "bbox": [
//         -79.8046875,
//         40.4771401,
//         -71.763627,
//         45.0239467
//       ],
//       "center": [
//         -75.4652471468304,
//         42.751210955038
//       ],
//       "geometry": {
//         "type": "Point",
//         "coordinates": [
//           -75.4652471468304,
//           42.751210955038
//         ]
//       },
//       "context": [
//         {
//           "id": "country.8940",
//           "mapbox_id": "dXJuOm1ieHBsYzpJdXc",
//           "wikidata": "Q30",
//           "short_code": "us",
//           "text": "United States"
//         }
//       ]
//     },
//     {
//       "id": "country.8940",
//       "type": "Feature",
//       "place_type": [
//         "country"
//       ],
//       "relevance": 1,
//       "properties": {
//         "mapbox_id": "dXJuOm1ieHBsYzpJdXc",
//         "wikidata": "Q30",
//         "short_code": "us"
//       },
//       "text": "United States",
//       "place_name": "United States",
//       "bbox": [
//         -179.9,
//         18.8164227,
//         -66.8847656,
//         71.420291
//       ],
//       "center": [
//         -97.9222112121185,
//         39.3812661305678
//       ],
//       "geometry": {
//         "type": "Point",
//         "coordinates": [
//           -97.9222112121185,
//           39.3812661305678
//         ]
//       }
//     }
//   ],
//   "attribution": "NOTICE: \u00A9 2023 Mapbox and its suppliers. All rights reserved. Use of this data is subject to the Mapbox Terms of Service (https://www.mapbox.com/about/maps/). This response and the information it contains may not be retained. POI(s) provided by Foursquare."
// }`
// 	var data map[string]interface{}

// 	if err := json.Unmarshal([]byte(bigData), &data); err != nil {
// 		fmt.Println("Invalid JSON structure:", err)
// 		return
// 	}

// 	features := data["features"].([]interface{})

// 	if len(features) > 0 {
// 		properties := features[0].(map[string]interface{})["properties"].(map[string]interface{})
// 		placeName := properties["place_name"].(string)
// 		fmt.Println("Place Name:", placeName)
// 	} else {
// 		fmt.Println("No features found")
// 	}
// }
