#GET ALL INDICES
GET /_cat/indices

#DESCRIBE an INDEX
GET foods-it-v2/_search
{}

###############################

#foodS Term CMS
GET foods-de-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "cmsId":{
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
GET filters-en/_doc/filter-foods
{
}

#FILTERS
GET filters-it/_search
{
}

#food
GET foods-de-v2/_search
{
  "query": {
    "match": {
      "roles.name": "ROLE_USER"
    }
  }
}


#food
GET foods-de-v2/_search
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

# "cmsId": "Tf6bnd4JkOgUkvsy7bMZg"

#food without legacyID
GET foods-en-v2/_search
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

GET foods-en-v2/_search
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

#food
GET foods-en-v2/

GET dishs-en-v2/

#food Terms cmsId
GET foods-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "terms":{
            "cmsid":["q7N94zX0PUmDuUrnptBhwP"],
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

#food TERMS legacyId, foodCategory
GET foods-it-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "terms":{
            "foodCategoryList.id":["8d0ad706-8e61-5f1a-abab-e525134ff66c"],
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

#food Terms title
GET dishs-en-v2/_search
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

#dish Term CMS
GET dishs-de-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "goal.name":{
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

#dish Term CMS
GET dishs-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "cmsId":{
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

#menuS Term shop CMS
GET menus-de-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "goal.legacyId":{
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

#foodS Term CMS
GET foods-en-v2/_search
{
  "query":{
    "bool":{
      "filter":[
        {
          "term":{
            "cmsId":{
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


#food
GET foods-en-v2/_search
{
  "query": {
    "match": {
      "cmsId": "q7N94zX0PUnB5105OYDEED"
    }
  }
}

#food
GET foods-de-v2/_search
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

#menuS
GET menus-de-v2/_search
{
  "query": {
    "match": {
      "goal.name": "Muskelaufbau"
    }
  }
}

#menuS PUBLISHED
GET menus-de-v2/_search
{
  "query": {
    "match": {
      "cmsId": "th4EfRogDnm3RdQv6Ozx3"
    }
  }
}

#menuS
GET menus-de-v2/_search
{
  "query": {
    "match": {
      "legacyId": 332
    }
  }
}


#dishS
GET dishs-en-v2/_search
{
  "sort": [
    {
      "popularity":"desc"
    }
  ]
}

#dishS
GET dishs-de-v2/_search
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
POST /foods-en-v2/_delete_by_query
{
  "query": { 
        "match_all": {}
    }
}

/* delete an entry in an index */
POST foods-de-v2/_delete_by_query
{
  "query": {
    "match": {
      "cmsId": "4NXTLIDZZRLFQhCJuOKIjv"
    }
  }
}

/* delete an entry in an index */
POST foods-en-v2/_delete_by_query
{
  "query": {
    "match": {
      "cmsId": "q7N94zX0PUmPbhYKnIB07l"
    }
  }
}

/** drop a field from index
POST durationfilters-de-v2/_delete_by_query
{
  "query": {
    "exists": { "field": "value" }
  }
}

/* drop an index */
DELETE /foods-de-v2

GET filters-en/_doc/filter-dishs
GET filters-en/_doc/filter-foods

GET filters-de/_doc/filter-dishs
GET filters-de/_doc/filter-foods

POST filters-de/_delete_by_query
{
  "query":{
    "match": {
      "foodCategories.id": "d3b257d3-6dee-569f-a69a-693aa67a17ab"
    }
  }
}

POST filters-en/_delete_by_query
{
  "query":{
    "match": {
      "foodCategories.id": "d3b257d3-6dee-569f-a69a-693aa67a17ab"
    }
  }
}


********************
RE-INDEX
*******************

/* 1. GET mappings for a current index that needs re-indexing*/
GET foods-de-v2
GET foods-de-v2/_mapping

/* 2. create a TEMP new index with same mappings */
PUT /foods-de-v3
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
      "cmsId": {
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
          "cmsId": {
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
              "cmsId": {
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
                  "cmsId": {
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
              "cmsId": {
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
      "foodCategoryList": {
        "properties": {
          "cmsId": {
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
      "foodDurationList": {
        "properties": {
          "cmsId": {
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
      "foodFoodTypeList": {
        "properties": {
          "cmsId": {
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
      "foodLevelOfDifficultyList": {
        "properties": {
          "cmsId": {
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
      "foodMealTypeList": {
        "properties": {
          "cmsId": {
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
      "foodPlateRuleList": {
        "properties": {
          "cmsId": {
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
      "foodVariantList": {
        "properties": {
          "cmsId": {
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

GET menus-de-v2/
GET foods-de-v3/_search
/* 3. reindex from the current to the temp */
POST /_reindex
{
  "source": {
    "index": "foods-de-v2"
  },
  "dest": {
    "index": "foods-de-v3"
  }
}

GET foods-de-v3/_search

/* 4. check data in the temp index, verify things are alright */

/* 5. delete the current index */
DELETE /foods-de-v2

/* 6. create the current index with same name and NEW mappings*/
PUT /foods-de-v2
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
  "cmsId": {
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
      "cmsId": {
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
          "cmsId": {
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
              "cmsId": {
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
          "cmsId": {
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
  "foodCategoryList": {
    "properties": {
      "cmsId": {
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
  "foodDurationList": {
    "properties": {
      "cmsId": {
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
  "foodFoodTypeList": {
    "properties": {
      "cmsId": {
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
  "foodLevelOfDifficultyList": {
    "properties": {
      "cmsId": {
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
  "foodMealTypeList": {
    "properties": {
      "cmsId": {
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
  "foodPlateRuleList": {
    "properties": {
      "cmsId": {
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
  "foodVariantList": {
    "properties": {
      "cmsId": {
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

GET menus-it-v2/_search
GET foods-de-v2/_search

/* 7. reindex from temp to current index */
POST /_reindex
{
  "source": {
    "index": "foods-de-v3"
  },
  "dest": {
    "index": "foods-de-v2"
  }
}

/* 8. verify things are alright with curr index */
GET foods-de-v2/_search

/* 9. delete the temp index */
DELETE /foods-de-v3



********************
CREATE INDEX
*******************
############
June 6, 2024, works fine on TEST
############
PUT /foods-de-v2
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
      "cmsId": {
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
          "cmsId": {
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
              "cmsId": {
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
                  "cmsId": {
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
              "cmsId": {
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
      "foodCategoryList": {
        "properties": {
          "cmsId": {
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
      "foodDurationList": {
        "properties": {
          "cmsId": {
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
      "foodFoodTypeList": {
        "properties": {
          "cmsId": {
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
      "foodLevelOfDifficultyList": {
        "properties": {
          "cmsId": {
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
      "foodMealTypeList": {
        "properties": {
          "cmsId": {
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
      "foodPlateRuleList": {
        "properties": {
          "cmsId": {
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
      "foodVariantList": {
        "properties": {
          "cmsId": {
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

PUT /foods-it-v2
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
      "cmsId": {
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
          "cmsId": {
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
              "cmsId": {
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
                  "cmsId": {
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
              "cmsId": {
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
      "foodCategoryList": {
        "properties": {
          "cmsId": {
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
      "foodDurationList": {
        "properties": {
          "cmsId": {
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
      "foodFoodTypeList": {
        "properties": {
          "cmsId": {
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
      "foodLevelOfDifficultyList": {
        "properties": {
          "cmsId": {
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
      "foodMealTypeList": {
        "properties": {
          "cmsId": {
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
      "foodPlateRuleList": {
        "properties": {
          "cmsId": {
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
      "foodVariantList": {
        "properties": {
          "cmsId": {
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

PUT /dishs-it-v2
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
        },
        "autocomplete_strict": {
          "filter": [
            "lowercase"
          ],
          "tokenizer": "autocomplete_strict"
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
        },
        "autocomplete_strict": {
          "token_chars": [
            "letter"
          ],
          "min_gram": "4",
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
      "categories": {
        "properties": {
          "categoryGroups": {
            "properties": {
              "cmsId": {
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
              "legacyId": {
                "type": "long"
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
              "name": {
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
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "complete": {
            "type": "boolean"
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
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
      "cmsId": {
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
        },
        "analyzer": "german"
      },
      "duration": {
        "properties": {
          "cmsId": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
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
          "legacyId": {
            "type": "long"
          },
          "value": {
            "type": "long"
          }
        }
      },
      "free": {
        "type": "boolean"
      },
      "goal": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "color": {
            "type": "keyword",
            "index": false
          },
          "complete": {
            "type": "boolean"
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
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
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "intensityLevel": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "complete": {
            "type": "boolean"
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
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
      "language": {
        "type": "keyword",
        "ignore_above": 256
      },
      "legacyId": {
        "type": "long"
      },
      "locale": {
        "type": "keyword",
        "ignore_above": 256
      },
      "metScore": {
        "type": "double"
      },
      "muscleGroup": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "complete": {
            "type": "boolean"
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
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
      "openTrainingEligible": {
        "type": "boolean"
      },
      "popularity": {
        "type": "long"
      },
      "published": {
        "type": "boolean"
      },
      "publishedAt": {
        "type": "date"
      },
      "roles": {
        "properties": {
          "name": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "shortDescription": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        },
        "analyzer": "german"
      },
      "slug": {
        "type": "keyword",
        "ignore_above": 256
      },
      "tags": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
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
          "name": {
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
      "tools": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "complete": {
            "type": "boolean"
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            },
            "analyzer": "autocomplete_strict"
          }
        }
      },
      "trainer": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "complete": {
            "type": "boolean"
          },
          "firstName": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "lastName": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            },
            "analyzer": "autocomplete_strict",
            "search_analyzer": "autocomplete_search"
          },
          "published": {
            "type": "boolean"
          }
        }
      },
      "updatedAt": {
        "type": "date"
      }
    }
  }
}

PUT /menus-it-v2
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
      "categories": {
        "properties": {
          "categoryGroups": {
            "properties": {
              "cmsId": {
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
              "legacyId": {
                "type": "long"
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
              "name": {
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
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "complete": {
            "type": "boolean"
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
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
      "cmsId": {
        "type": "keyword",
        "ignore_above": 256
      },
      "configuration": {
        "properties": {
          "maxTrainingDays": {
            "type": "integer",
            "index": false
          },
          "minTrainingDays": {
            "type": "integer",
            "index": false
          },
          "optimalTrainingDays": {
            "type": "integer",
            "index": false
          },
          "preview": {
            "type": "boolean"
          },
          "replacedishAllowed": {
            "type": "boolean",
            "index": false
          },
          "dishsRecommended": {
            "type": "long"
          }
        }
      },
      "continuous": {
        "type": "boolean"
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
        },
        "analyzer": "german"
      },
      "duration": {
        "properties": {
          "cmsId": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
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
          "legacyId": {
            "type": "long"
          },
          "value": {
            "type": "long"
          }
        }
      },
      "gender": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "name": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "goal": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "color": {
            "type": "keyword",
            "index": false
          },
          "complete": {
            "type": "boolean"
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
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
      "id": {
        "type": "keyword",
        "ignore_above": 256
      },
      "intensityLevel": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "complete": {
            "type": "boolean"
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
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
      "language": {
        "type": "keyword",
        "ignore_above": 256
      },
      "legacyId": {
        "type": "long"
      },
      "locale": {
        "type": "keyword",
        "ignore_above": 256
      },
      "popularity": {
        "type": "long"
      },
      "preview": {
        "type": "boolean",
        "index": false
      },
      "previewVideoId": {
        "type": "keyword",
        "index": false
      },
      "primaryTerm": {
        "type": "long"
      },
      "published": {
        "type": "boolean"
      },
      "publishedAt": {
        "type": "date"
      },
      "punchline": {
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
      "roles": {
        "properties": {
          "name": {
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
      "seqNo": {
        "type": "long"
      },
      "shopItems": {
        "properties": {
          "cmsId": {
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
          "title": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "url": {
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
      "shortDescription": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        },
        "analyzer": "german"
      },
      "slug": {
        "type": "keyword",
        "ignore_above": 256
      },
      "subtitle": {
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
      "tags": {
        "properties": {
          "cmsId": {
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
          "legacyId": {
            "type": "long"
          },
          "name": {
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
      "trainer": {
        "properties": {
          "cmsId": {
            "type": "keyword",
            "ignore_above": 256
          },
          "complete": {
            "type": "boolean"
          },
          "firstName": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "id": {
            "type": "keyword",
            "ignore_above": 256
          },
          "lastName": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "legacyId": {
            "type": "long"
          },
          "name": {
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
      "updatedAt": {
        "type": "date"
      },
      "version": {
        "type": "long"
      }
    }
  }
}
