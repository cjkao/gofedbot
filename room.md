```json
{
  "_id": "22nytdn4QqKvSk2Av",              // Random.id()
  "t": "p",                                // String           Room Type: c = chanel, d = direct, p (change to g) = group, v = visitor
  "ts": new Date(1432937851208),           // Date             Room Creation Timestamp
  "name": "general",                       // String           Room Name (t:d -> empty)
  "fname": full name,(display)
  "lm": new Date(1432937851208),           // Date             Last Message Timestamp
  "msgs": 2345,                            // Integer          Messages Counter
  "cl": true,                              // Boolean          If users can leave room
  "ro": false,                             // Boolean          Read Only
  "usernames": [                           // Array(String)    Room Users
      "username1",                         // String           User Username
      "username2",
      "username3"
  ],
  "u": {                                   // Object           Owner User
    "_id": "CABrXSHpDqbZXGkYR",            // Random.id()      User Id
    "username": "john"                     // String           User Username
  },
  "customFields": {                        // Object           User defined custom fields (for t:c and t:p only)
     "userDefinedField": "userValue",      //                  User defined field example
     "userDefinedField2": true,            //                  User defined field example
     "userObject3": {                      //                  User defined field example
        "a": "hello",
        "b": "lalala"
     }
  }
}
```