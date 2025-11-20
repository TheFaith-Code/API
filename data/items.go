package data

import "API/models"

// In-memory DB
var Items = []models.Item{
    {ID: 1, Name: "Alpha", Description: "First demo item"},
    {ID: 2, Name: "Beta", Description: "Second demo item"},
    {ID: 3, Name: "Gamma", Description: "Third demo item"},
}
