package integration

import (
	"cms-backend/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TODO: Import required packages for:
// - JSON handling
// - HTTP testing
// - Your application models
// - Testing package

/*
MEDIA INTEGRATION TESTS

These tests verify the complete flow of media operations through the API.
Each test should:
1. Start with a clean database state
2. Perform API operations
3. Verify the responses
4. Check database state if needed
*/

func TestMediaIntegration(t *testing.T) {
	//TODO: Implement TestMediaIntegration
    // STEP 1: Clear Database
    // - Clear all tables before starting tests
	clearTables()

	t.Run("Create Media", func(t *testing.T) {
		//TODO: Implement test logic
        // STEP 1: Prepare Test Data
        // - Create JSON body with:
        //   * URL (e.g., "http://example.com/test.jpg")
        //   * Type (e.g., "image")
		body := `{
			"url": "http://example.com/test.jpg",
			"type": "image"
		}`

		// STEP 2: Create HTTP Request
		// - Create POST request to /api/v1/media
		// - Set Content-Type header
		// - Add request body	
		req := httptest.NewRequest("POST", "/api/v1/media", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// STEP 3: Execute Request
		// - Create response recorder
		// - Send request through router
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Fatalf("Expected status 201, got %d: %s", w.Code, w.Body.String())
		}

		// STEP 4: Verify Response
		// - Check status code (should be 201 Created)
		// - Parse response JSON
		// - Verify media properties (URL, type)
		// - Handle any parsing errors
		var response models.Media
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		if response.URL != "http://example.com/test.jpg" {
			t.Errorf("Expected URL 'http://example.com/test.jpg', got %s", response.URL)
		}
	})

    t.Run("Get All Media", func(t *testing.T) {
		//TODO: Implement test logic
        // STEP 1: Setup Test Data
        // - Create test media entries if needed
        
        // STEP 2: Create HTTP Request
        // - Create GET request to /api/v1/media
        
        // STEP 3: Execute Request
        // - Create response recorder
        // - Send request through router
        
        // STEP 4: Verify Response
        // - Check status code (should be 200 OK)
        // - Parse response JSON array
        // - Verify media list properties
        // - Check number of items
    })

    // TODO: Additional test cases to consider:
    // - Get single media by ID
    // - Create media with invalid data
    // - Delete media
    // - Create media with duplicate URL
}

/*
TESTING HINTS:
1. Request Creation:
   - Use httptest.NewRequest for creating requests
   - Remember to set Content-Type for POST requests
   - Use strings.NewReader for request bodies

2. Response Handling:
   - Use httptest.NewRecorder for capturing responses
   - Parse JSON responses carefully
   - Check both status codes and response bodies

3. Test Data:
   - Use meaningful test data
   - Clean up between tests
   - Consider edge cases

4. Error Cases:
   - Test invalid inputs
   - Test missing required fields
   - Test invalid content types
*/