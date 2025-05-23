/*
 * Waiting List Api
 *
 * Consulting Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xtodorovic@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package consulting

import (
	"github.com/gin-gonic/gin"
)

type RequestsListAPI interface {


    // GetRequestsListEntries Get /api/requests-list/:requestId/entries
    // Provides the requests list 
     GetRequestsListEntries(c *gin.Context)

}