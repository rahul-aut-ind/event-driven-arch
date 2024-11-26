#GET ALL INDICES
GET /_cat/indices

#DESCRIBE an INDEX
GET menu-it-v2/_search
{}

###############################

#menu Term EXT
GET menu-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "extId":{
              "value":"q7N94zX0PUmPiFQiidT1f1",
              "boost":1.0
            }
          }
        },
        {
          "term":{
            "published":{
              "value":"true",
              "boost":1.0
            }
          }
        }
      ]
    }
  }
}

#FILTERS
GET filters-en/_doc/filter-menu
{
}

#FILTERS
GET filters-it/_search
{
}

#menu
GET menu-en-v2/_search
{
  "query": {
    "match": {
      "roles.name": "ROLE_USER"
    }
  }
}


#menu
GET menu-en-v2/_search
{
  "query": {
    "constant_score": {
      "filter": {
        "exists": {
          "field": "popularity"
        }
      }
    }
  }
}

# "extId": "Tf6bnd4JkOgUkvsy7bMZg"

#menu without legacyID
GET menu-en-v2/_search
{
  "query": {
    "bool": {
      "must_not": {
        "exists": {
          "field": "roles"
        }
      }
    }
  }
}

GET menu-en-v2/_search
{
  "query": {
    "range": {
      "calories": {
                "gte" : 500,
                "lte" : 2000
            }
    }
  }
}

#menu
GET menu-en-v2/

GET audio-en-v2/

#menu Terms extId
GET menu-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "terms":{
            "extId":["q7N94zX0PUmDuUrnptBhwP"],
            "boost":1.0
          }
        },
        {
          "term":{
            "published":{
              "value":"true",
              "boost":1.0
            }
          }
        }
      ]
    }
  }
}

#menu TERMS legacyId, menuCategory
GET menu-it-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "terms":{
            "menuCategoryList.id":["8d0ad706-8e61-5f1a-abab-e525134ff66c"],
            "boost":1.0
          }
        },
        {
          "term":{
            "published":{
              "value":"true",
              "boost":1.0
            }
          }
        }
      ]
    }
  }
}

#menu Terms title
GET audio-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "terms":{
            "categories.id" : [
            "Beginner"
          ],
            "boost":1.0
          }
        }
      ]
    }
  }
}

#AUDIO Term EXT
GET audio-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "genre.name":{
              "value":"Straffen",
              "boost":1.0
            }
          }
        },
        {
          "term":{
            "published":{
              "value":"true",
              "boost":1.0
            }
          }
        }
      ]
    }
  }
}

#AUDIO Term EXT
GET audio-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "extId":{
              "value":"rgTWaEzhbk7qUnfV43Z7qf",
              "boost":1.0
            }
          }
        },
        {
          "term":{
            "published":{
              "value":"true",
              "boost":1.0
            }
          }
        }
      ]
    }
  }
}

#playlists Term shop EXT
GET playlists-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "genre.legacyId":{
              "value":"13",
              "boost":1.0
            }
          }
        },
        {
          "term":{
            "published":{
              "value":"true",
              "boost":1.0
            }
          }
        }
      ]
    }
  }
}

#menu Term EXT
GET menu-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "extId":{
              "value":"q7N94zX0PUmDuUrnptBhwP",
              "boost":1.0
            }
          }
        },
        {
          "term":{
            "published":{
              "value":"true",
              "boost":1.0
            }
          }
        }
      ]
    }
  }
}


#menu
GET menu-en-v2/_search
{
  "query": {
    "match": {
      "extId": "q7N94zX0PUnB5105OYDEED"
    }
  }
}

#menu
GET menu-en-v2/_search
{
  "query": {
    "constant_score": {
      "filter" : {
        "exists" : {
           "field" : "legacyId"
        }
      }
    }
  }
}

#playlists
GET playlists-en-v2/_search
{
  "query": {
    "match": {
      "genre.name": "Muskelaufbau"
    }
  }
}

#playlists PUBLISHED
GET playlists-en-v2/_search
{
  "query": {
    "match": {
      "extId": "th4EfRogDnm3RdQv6Ozx3"
    }
  }
}

#playlists
GET playlists-en-v2/_search
{
  "query": {
    "match": {
      "legacyId": 332
    }
  }
}


#AUDIOS
GET audio-en-v2/_search
{
  "sort": [
    {
      "popularity":"desc"
    }
  ]
}

#AUDIOS
GET audio-en-v2/_search
{
  "query": {
    "match": {
      "legacyId" : 3293
    }
  }
}

###############################
###############################
###############################
**************
DANGER ZONE
**************

/*** drops all document form index type
POST /menu-en-v2/_delete_by_query
{
  "query": { 
        "match_all": {}
    }
}

/* delete an entry in an index */
POST menu-en-v2/_delete_by_query
{
  "query": {
    "match": {
      "extId": "4NXTLIDZZRLFQhCJuOKIjv"
    }
  }
}

/* delete an entry in an index */
POST menu-en-v2/_delete_by_query
{
  "query": {
    "match": {
      "extId": "q7N94zX0PUmPbhYKnIB07l"
    }
  }
}

/** drop a field from index
POST userplaylist-en-v2/_delete_by_query
{
  "query": {
    "exists": { "field": "value" }
  }
}

/* drop an index */
DELETE /menu-en-v2

GET filters-en/_doc/filter-audio
GET filters-en/_doc/filter-menu

GET filters-de/_doc/filter-audio
GET filters-de/_doc/filter-menu

POST filters-de/_delete_by_query
{
  "query":{
    "match": {
      "menuCategories.id": "d3b257d3-6dee-569f-a69a-693aa67a17ab"
    }
  }
}

POST filters-en/_delete_by_query
{
  "query":{
    "match": {
      "menuCategories.id": "d3b257d3-6dee-569f-a69a-693aa67a17ab"
    }
  }
}


********************
RE-INDEX
*******************

/* 1. GET mappings for a current index that needs re-indexing*/
GET menu-en-v2
GET menu-en-v2/_mapping

/* 2. create a TEMP new index with same mappings */
PUT /menu-en-v3
{
  "settings": {
    "number_of_shards": 2,
    "analysis": {
      "analyzer": {
        "autocomplete": {
          "filter": [
            "lowercase"
          ],
          "tokenizer": "autocomplete"
        },
        "autocomplete_search": {
          "tokenizer": "lowercase"
        }
      },
      "tokenizer": {
        "autocomplete": {
          "token_chars": [
            "letter"
          ],
          "min_gram": "2",
          "type": "edge_ngram",
          "max_gram": "10"
        }
      }
    }
  },
  "mappings": {
    "_source": {
      "enabled": true
    },
    "properties": {
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "legacyId": {
        "type": "long"
      },
      "locale": {
        "type": "text"
      },
      "calories": {
        "type": "float"
      },
      "carbohydrateDietaryFiber": {
        "type": "float"
      },
      "description": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "descriptionShort": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "freebie": {
        "type": "boolean"
      },
      "ingredientAmountList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "id": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "ingredient": {
            "properties": {
              "extId": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "gramsPerCup": {
                "type": "float"
              },
              "id": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "labelPlural": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "labelSingular": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "locale": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "published": {
                "type": "boolean"
              },
              "supermarketAisle": {
                "properties": {
                  "extId": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "id": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "label": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "locale": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  }
                }
              }
            }
          },
          "ingredientUnit": {
            "properties": {
              "extId": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "id": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "label": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "locale": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              }
            }
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "multipliable": {
            "type": "boolean"
          },
          "published": {
            "type": "boolean"
          },
          "quantity": {
            "type": "long"
          }
        }
      },
      "protein": {
        "type": "float"
      },
      "published": {
        "type": "boolean"
      },
      "menuCategoryList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuDurationList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuFoodTypeList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuLevelOfDifficultyList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuMealTypeList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuPlateRuleList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuVariantList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "servingAmount": {
        "type": "long"
      },
      "subtitle": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "tip": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "title": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        },
        "analyzer": "autocomplete",
        "search_analyzer": "autocomplete_search"
      },
      "totalCarbohydrate": {
        "type": "float"
      },
      "totalFat": {
        "type": "float"
      }
    }
  }
}

GET menu-en-v3/
GET menu-en-v3/_search
/* 3. reindex from the current to the temp */
POST /_reindex
{
  "source": {
    "index": "menu-en-v2"
  },
  "dest": {
    "index": "menu-en-v3"
  }
}

GET menu-en-v3/_search

/* 4. check data in the temp index, verify things are alright */

/* 5. delete the current index */
DELETE /menu-en-v2

/* 6. create the current index with same name and NEW mappings*/
PUT /menu-en-v2
{
  "settings": {
    "number_of_shards": 2,
    "analysis": {
      "analyzer": {
        "autocomplete": {
          "filter": [
            "lowercase"
          ],
          "tokenizer": "autocomplete"
        },
        "autocomplete_search": {
          "tokenizer": "lowercase"
        }
      },
      "tokenizer": {
        "autocomplete": {
          "token_chars": [
            "letter"
          ],
          "min_gram": "2",
          "type": "edge_ngram",
          "max_gram": "10"
        }
      }
    }
  },
  "mappings": {
    "_source": {
      "enabled": true
    },
    "properties": {
  "calories": {
    "type": "float"
  },
  "carbohydrateDietaryFiber": {
    "type": "float"
  },
  "extId": {
    "type": "keyword",
    "ignore_above": 256
  },
  "createdAt": {
    "type": "date"
  },
  "description": {
    "type": "text",
    "fields": {
      "keyword": {
        "type": "keyword",
        "ignore_above": 256
      }
    }
  },
  "descriptionShort": {
    "type": "text",
    "fields": {
      "keyword": {
        "type": "keyword",
        "ignore_above": 256
      }
    }
  },
  "freebie": {
    "type": "boolean"
  },
  "id": {
    "type": "keyword",
    "ignore_above": 256
  },
  "ingredientAmountList": {
    "properties": {
      "extId": {
        "type": "keyword",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "id": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "ingredient": {
        "properties": {
          "extId": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "gramsPerCup": {
            "type": "float"
          },
          "id": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "labelPlural": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "labelSingular": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          },
          "supermarketAisle": {
            "properties": {
              "extId": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "id": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "label": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "locale": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              }
            }
          }
        }
      },
      "ingredientUnit": {
        "properties": {
          "extId": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "id": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          }
        }
      },
      "label": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "locale": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "multipliable": {
        "type": "boolean"
      },
      "published": {
        "type": "boolean"
      },
      "quantity": {
        "type": "long"
      }
    }
  },
  "legacyId": {
    "type": "long"
  },
  "locale": {
    "type": "text"
  },
  "protein": {
    "type": "float"
  },
  "published": {
    "type": "boolean"
  },
  "menuCategoryList": {
    "properties": {
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "label": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "locale": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "published": {
        "type": "boolean"
      }
    }
  },
  "menuDurationList": {
    "properties": {
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "label": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "locale": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "published": {
        "type": "boolean"
      }
    }
  },
  "menuFoodTypeList": {
    "properties": {
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "label": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "locale": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "published": {
        "type": "boolean"
      }
    }
  },
  "menuLevelOfDifficultyList": {
    "properties": {
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "label": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "locale": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "published": {
        "type": "boolean"
      }
    }
  },
  "roles": {
    "properties": {
      "name": {
        "type": "keyword",
        "ignore_above": 256
      }
    }
  },
  "menuMealTypeList": {
    "properties": {
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "label": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "locale": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "published": {
        "type": "boolean"
      }
    }
  },
  "menuPlateRuleList": {
    "properties": {
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "label": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "locale": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "published": {
        "type": "boolean"
      }
    }
  },
  "menuVariantList": {
    "properties": {
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "label": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "locale": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "published": {
        "type": "boolean"
      }
    }
  },
  "servingAmount": {
    "type": "long"
  },
  "subtitle": {
    "type": "text",
    "fields": {
      "keyword": {
        "type": "keyword",
        "ignore_above": 256
      }
    }
  },
  "tip": {
    "type": "text",
    "fields": {
      "keyword": {
        "type": "keyword",
        "ignore_above": 256
      }
    }
  },
  "title": {
    "type": "text",
    "fields": {
      "keyword": {
        "type": "keyword",
        "ignore_above": 256
      }
    },
    "analyzer": "autocomplete",
    "search_analyzer": "autocomplete_search"
  },
  "totalCarbohydrate": {
    "type": "float"
  },
  "totalFat": {
    "type": "float"
  },
  "updatedAt": {
    "type": "date"
  }
}
  }
}

GET menu-en-v2/
GET menu-en-v2/_search

/* 7. reindex from temp to current index */
POST /_reindex
{
  "source": {
    "index": "menu-en-v3"
  },
  "dest": {
    "index": "menu-en-v2"
  }
}

/* 8. verify things are alright with curr index */
GET menu-en-v2/_search

/* 9. delete the temp index */
DELETE /menu-en-v3



********************
CREATE INDEX
*******************
############
June 6, 2024, works fine on TEST
############
PUT /menu-en-v2
{
  "settings": {
    "number_of_shards": 2,
    "analysis": {
      "analyzer": {
        "autocomplete": {
          "filter": [
            "lowercase"
          ],
          "tokenizer": "autocomplete"
        },
        "autocomplete_search": {
          "tokenizer": "lowercase"
        }
      },
      "tokenizer": {
        "autocomplete": {
          "token_chars": [
            "letter"
          ],
          "min_gram": "2",
          "type": "edge_ngram",
          "max_gram": "10"
        }
      }
    }
  },
  "mappings": {
    "_source": {
      "enabled": true
    },
    "properties": {
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "legacyId": {
        "type": "long"
      },
      "locale": {
        "type": "text"
      },
      "calories": {
        "type": "float"
      },
      "carbohydrateDietaryFiber": {
        "type": "float"
      },
      "description": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "descriptionShort": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "freebie": {
        "type": "boolean"
      },
      "ingredientAmountList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "id": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "ingredient": {
            "properties": {
              "extId": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "gramsPerCup": {
                "type": "float"
              },
              "id": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "labelPlural": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "labelSingular": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "locale": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "published": {
                "type": "boolean"
              },
              "supermarketAisle": {
                "properties": {
                  "extId": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "id": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "label": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "locale": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  }
                }
              }
            }
          },
          "ingredientUnit": {
            "properties": {
              "extId": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "id": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "label": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "locale": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              }
            }
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "multipliable": {
            "type": "boolean"
          },
          "published": {
            "type": "boolean"
          },
          "quantity": {
            "type": "long"
          }
        }
      },
      "protein": {
        "type": "float"
      },
      "published": {
        "type": "boolean"
      },
      "menuCategoryList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuDurationList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuFoodTypeList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuLevelOfDifficultyList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuMealTypeList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuPlateRuleList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuVariantList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "servingAmount": {
        "type": "long"
      },
      "subtitle": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "tip": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "title": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        },
        "analyzer": "autocomplete",
        "search_analyzer": "autocomplete_search"
      },
      "totalCarbohydrate": {
        "type": "float"
      },
      "totalFat": {
        "type": "float"
      }
    }
  }
}

PUT /menu-it-v2
{
  "settings": {
    "number_of_shards": 2,
    "analysis": {
      "analyzer": {
        "autocomplete": {
          "filter": [
            "lowercase"
          ],
          "tokenizer": "autocomplete"
        },
        "autocomplete_search": {
          "tokenizer": "lowercase"
        }
      },
      "tokenizer": {
        "autocomplete": {
          "token_chars": [
            "letter"
          ],
          "min_gram": "2",
          "type": "edge_ngram",
          "max_gram": "10"
        }
      }
    }
  },
  "mappings": {
    "_source": {
      "enabled": true
    },
    "properties": {
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "extId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "legacyId": {
        "type": "long"
      },
      "locale": {
        "type": "text"
      },
      "calories": {
        "type": "float"
      },
      "carbohydrateDietaryFiber": {
        "type": "float"
      },
      "description": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "descriptionShort": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "freebie": {
        "type": "boolean"
      },
      "ingredientAmountList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "id": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "ingredient": {
            "properties": {
              "extId": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "gramsPerCup": {
                "type": "float"
              },
              "id": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "labelPlural": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "labelSingular": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "locale": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "published": {
                "type": "boolean"
              },
              "supermarketAisle": {
                "properties": {
                  "extId": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "id": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "label": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "locale": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  }
                }
              }
            }
          },
          "ingredientUnit": {
            "properties": {
              "extId": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "id": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "label": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "locale": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              }
            }
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "multipliable": {
            "type": "boolean"
          },
          "published": {
            "type": "boolean"
          },
          "quantity": {
            "type": "long"
          }
        }
      },
      "protein": {
        "type": "float"
      },
      "published": {
        "type": "boolean"
      },
      "menuCategoryList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuDurationList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuFoodTypeList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuLevelOfDifficultyList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuMealTypeList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuPlateRuleList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "menuVariantList": {
        "properties": {
          "extId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "label": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "locale": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "servingAmount": {
        "type": "long"
      },
      "subtitle": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "tip": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "title": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        },
        "analyzer": "autocomplete",
        "search_analyzer": "autocomplete_search"
      },
      "totalCarbohydrate": {
        "type": "float"
      },
      "totalFat": {
        "type": "float"
      }
    }
  }
}

