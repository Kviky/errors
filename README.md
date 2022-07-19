# Ahoylog errors package

Ahoylog package used together with [go-swagger](https://github.com/go-swagger/go-swagger) to produce standardized set of errors as [ProblemDetails](https://tools.ietf.org/html/rfc7807). 

## List of errors

### HTTP **400**
| Title | Detail | Status | Code | Instance |
| --- | --- | --- | --- | --- | 
| AlreadyExists | The requested resource already exists! | 400 | badRequest | client |
| BadRequest | There was a problem with the request! | 400 | badRequest | client |
| CharterHasListings | Charter cannot be deleted, because it still has some active listings! | 400 | badRequest | client |
| CharterNotCreated | There was a problem to create charter profile! | 400 | badRequest | client |
| FileExistsAlready | The requested resource already exists! | 400 | badRequest | client |
| FileNotCreated | There was a problem to create file! | 400 | badRequest | export |
| InvalidBodyParam | The HTTP request contains an unsupported body parameter! | 400 | badRequest | client |
| InvalidDates | The requested dates are invalid! | 400 | badRequest | client |
| InvalidHeaderParam | The HTTP request contains an unsupported header parameter! | 400 | badRequest | client |
| InvalidMsgFormat | The HTTP request has an invalid format! | 400 | badRequest | client |
| ImageInvalid | File must be a valid image - image/jpeg, image/jpg, image/png! | 400 | badRequest | client |
| ImageNotDeleted | There was a problem to delete image! | 400 | badRequest | image |
| ImageNotUploaded | There was a problem to upload image! | 400 | badRequest | image |
| InactiveListing | Listing %v is not in the active state! | 400 | badRequest | client |
| InvalidOwnerListing | Charter doesn't own the listing %v! | 400 | badRequest | client |
| InvalidQueryParam | The HTTP request contains an unsupported query parameter in the URI! | 400 | badRequest | client |
| InvalidPathParam | The HTTP request contains an unsupported path parameter in the URI! | 400 | badRequest | client |
| ListingNotCreated | There was a problem to create listing! | 400 | badRequest | client |
| LocationNotCreated | There was a problem to create location! | 400 | badRequest | client |
| MandatoryParamIncorrect | Mandatory parameter has semantically incorrect value! | 400 | badRequest | client |
| MandatoryParamMissing | Parameter which is defined as mandatory is missing! | 400 | badRequest | client |
| NameAlreadyTaken | Requested name is already taken! Please, specify another name. | 400 | badRequest | client |
| OffersEnded | Available number of the offers ended for today! | 400 | badRequest | client |
| OffersMaxListings | Maximum limit of %v listings is reached. Please, reduce number of listings in offer! | 400 | badRequest | client |
| PortAlreadyExists | Requested port/marina name already exists for this country and city! | 400 | badRequest | client |
| ReservationNotCreated | There was a problem to create reservation! | 400 | badRequest | client |

### HTTP **401**
| Title | Detail | Status | Code | Instance |
| --- | --- | --- | --- | --- | 
| InvalidAuthToken | Authorization token is invalid! | 401 | unauthorized | client |
| MissingAuthToken | Authorization token is missing! | 401 | unauthorized | client |
| UnauthorizedAccess | The request doesn't have permissions to access resources! | 401 | unauthorized | api |

### HTTP **403**
| Title | Detail | Status | Code | Instance |
| --- | --- | --- | --- | --- | 
| ForbiddenAction | You don't have a permission to make this action! | 403 | forbidden | client |
| ForbiddenResource | You don't have a permission to access this resource! | 403 | forbidden | client |
| ForbiddenUpload | This accound doesn't have permission to upload images! | 403 | forbidden | client |

### HTTP **404**
| Title | Detail | Status | Code | Instance |
| --- | --- | --- | --- | --- | 
| CharterNotFound | The charter indicated in the request does not exist! | 404 | notFound | client |
| ListingNotFound | The listing indicated in the request does not exist! | 404 | notFound | client |
| LocationNotFound | The location indicated in the request does not exist! | 404 | notFound | client |
| ReservationNotFound | Requested reservation does not exist! | 404 | notFound | client |
| ResourceNotFound | Requested resource does not exist! | 404 | notFound | client |
| UserNotFound | The user indicated in the request does not exist! | 404 | notFound | client |
| UsersNotFound | Requested users does not exist! | 404 | notFound | client |

### HTTP **405**
| Title | Detail | Status | Code | Instance |
| --- | --- | --- | --- | --- | 
| MethodNotAllowed | Requested method is not allowed. Check the response header `Allow` for allowed methods! | 405 | methodNotAllowed | client |

### HTTP **429**
| Title | Detail | Status | Code | Instance |
| --- | --- | --- | --- | --- | 
| CongestionRisk | The request is rejected due to excessive traffic. If continued over time, may lead to an overload situation. | 429 | tooManyRequests | client |

### HTTP **500**
| Title | Detail | Status | Code | Instance |
| --- | --- | --- | --- | --- | 
| UnspecifiedFailure | The request is rejected due to unspecified reason at the system! | 500 | internalServerError | api |
| SystemFailure | We are sorry, but there is an internal problem with the application! | 500 | internalServerError | api |

### HTTP **503**
| Title | Detail | Status | Code | Instance |
| --- | --- | --- | --- | --- | 
| ServiceUnavailable | The service experiences congestion and performs overload control. It does not allow the request to be processed. | 503 | serviceUnavailable | api |

### HTTP **504**
| Title | Detail | Status | Code | Instance |
| --- | --- | --- | --- | --- | 
| GatewayTimeout | The request is rejected due a request that has timed out at the HTTP client. | 504 | gatewayTimeout | api |



