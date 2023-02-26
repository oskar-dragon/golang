package maps

import (
	"testing"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want error %q", got, want)
	}
	
}

func TestMaps(t *testing.T) {
	dictionary := Dictionary{"test": "definition of a test"}

	t.Run("gets a word", func(t *testing.T) {
		got,_ := dictionary.Search("test")
		want := "definition of a test"

		assertStrings(t, got, want)
	})

	t.Run("returns error when word does not exist", func(t *testing.T) {
		_,err := dictionary.Search("test2")
		want := ErrWordNotFound

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		definition := "this is just a test"
		word := "test"
		
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		definition := "this is just a test"
		word := "test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)
		
		assertError(t, err, ErrWordAlreadyExist)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T){
		definition := "this is just a test"
		word := "test"
		dictionary := Dictionary{word: definition}
		newDefinition := "this is updated definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update word that does not exist", func(t *testing.T) {
		dictionary := make(Dictionary)
		word := "test"
		definition := "this is definition"

		err := dictionary.Update(word, definition)

		assertError(t,err, ErrWordCantBeUpdated)
	})
	
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is definition"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrWordNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}