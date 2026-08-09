package maps

import (
	"testing"
)

const (
	testKey              = "test"
	testDefinition       = "test is process to verify a fact"
	testUpdateDefinition = "test is good"
	unknownKey           = "unknown"
)

func TestSearch(t *testing.T) {

	t.Run("Test Search 1 : Searching for existing word", func(t *testing.T) {
		dictionary := Dictionary{testKey: testDefinition}

		got, err := dictionary.Search(testKey)
		wanted := testDefinition
		assertNotError(t, err)
		assertString(t, got, wanted)
	})

	t.Run("Test Search 2 : Searching an unknown word", func(t *testing.T) {
		dictionary := Dictionary{}

		_, err := dictionary.Search(unknownKey)
		assertError(t, err, missingWordError)
	})
}

func assertError(t testing.TB, got, wanted error) {

	t.Helper()
	if got == nil {
		t.Errorf("An error should be raisde")
	}

	assertString(t, got.Error(), wanted.Error())
}

func assertNotError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("An error has been raised : %v", got)
	}
}

func assertString(t testing.TB, got, wanted string) {
	t.Helper()

	if got != wanted {
		t.Errorf("got %s is not equal to wanted %s", got, wanted)
	}
}

func TestAdd(t *testing.T) {

	t.Run("Test Add 1 : Adding new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Add(testKey, testDefinition)
		assertNotError(t, err)
		got, err := dictionary.Search(testKey)

		assertNotError(t, err)
		assertString(t, got, testDefinition)
	})

	t.Run("Test Add 2 : Add an existing word", func(t *testing.T) {
		dictionary := Dictionary{testKey: testDefinition}

		err := dictionary.Add(testKey, testDefinition)
		assertError(t, err, duplicateWordError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Update 1 : Update a existing word", func(t *testing.T) {
		dictionary := Dictionary{testKey: testDefinition}
		err := dictionary.Update(testKey, testUpdateDefinition)

		assertNotError(t, err)

		got, err := dictionary.Search(testKey)

		assertNotError(t, err)
		assertString(t, got, testUpdateDefinition)
	})

	t.Run("Test Update 2 : Update a non existing word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update(testKey, testDefinition)
		assertError(t, err, updateOperationError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Delete 1 : Delete an existing world", func(t *testing.T) {
		dictionary := Dictionary{testKey: testDefinition}

		err := dictionary.Delete(testKey)

		assertNotError(t, err)

		got, err := dictionary.Search(testKey)

		assertError(t, err, missingWordError)
		assertString(t, got, "")
	})

	t.Run("Test Delete 2 : Delete an non existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete(testKey)

		assertError(t, err, deleteOperationError)
	})
}
