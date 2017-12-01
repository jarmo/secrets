package vault

import (
  "testing"
  "fmt"
  "os"
  "io/ioutil"
  "github.com/satori/go.uuid"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/vault/storage"
)

func TestList(t *testing.T) {
  vaultPath := prepareVault(t)
  defer os.Remove(vaultPath)

  filter := ""
  listedSecrets := List(storage.Read(password(), vaultPath), filter)

  expectedListedSecrets := `[
[0da52a01-302d-4e2f-8200-a4d4226699af]
secret-1-name
secret-1-value 
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-name
secret-2-value 
[1f947c00-211c-4216-865e-1ca1fcbce693]
secret-3-name
secret-3-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %s, but got %s", expectedListedSecrets, listedSecrets))
  }
}

func TestList_WithFilter(t *testing.T) {
  vaultPath := prepareVault(t)
  defer os.Remove(vaultPath)

  filter := "secret-2"
  listedSecrets := List(storage.Read(password(), vaultPath), filter)

  expectedListedSecrets := `[
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-name
secret-2-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %s, but got %s", expectedListedSecrets, listedSecrets))
  }
}

func TestAdd(t *testing.T) {
  vaultPath := prepareVault(t)
  defer os.Remove(vaultPath)

  addedSecret := Add("secret-4-name", "secret-4-value", vaultPath, password())

  filter := ""
  listedSecrets := List(storage.Read(password(), vaultPath), filter)

  expectedListedSecrets := fmt.Sprintf(`[
[0da52a01-302d-4e2f-8200-a4d4226699af]
secret-1-name
secret-1-value 
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-name
secret-2-value 
[1f947c00-211c-4216-865e-1ca1fcbce693]
secret-3-name
secret-3-value 
[%s]
secret-4-name
secret-4-value]`, addedSecret.Id)

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %v, but got %v", expectedListedSecrets, listedSecrets))
  }
}

func TestDelete(t *testing.T) {
  vaultPath := prepareVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-c70b-4a81-87db-7839d16f0176")
  if deletedSecret, err := Delete(id, vaultPath, password()); err != nil {
    t.Fatal(err)
  } else {
    expectedDeletedSecret := secret.Secret{id, "secret-2-name", "secret-2-value"}
    if fmt.Sprintf("%v", expectedDeletedSecret) != fmt.Sprintf("%v", deletedSecret) {
      t.Fatal(fmt.Sprintf("Expected deleted secret to be %s, but got %s", expectedDeletedSecret, deletedSecret))
    }
  }

  filter := ""
  listedSecrets := List(storage.Read(password(), vaultPath), filter)

  expectedListedSecrets := `[
[0da52a01-302d-4e2f-8200-a4d4226699af]
secret-1-name
secret-1-value 
[1f947c00-211c-4216-865e-1ca1fcbce693]
secret-3-name
secret-3-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %v, but got %v", expectedListedSecrets, listedSecrets))
  }
}

func TestDelete_NonExistingId(t *testing.T) {
  vaultPath := prepareVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-0000-0000-87db-7839d16f0176")
  expectedError := "Secret by specified id not found!"
  if deletedSecret, err := Delete(id, vaultPath, password()); err.Error() != expectedError {
    t.Fatal("Expected to return an error %s, but got %s", expectedError, err.Error())
  } else if deletedSecret != nil {
    t.Fatal(fmt.Sprintf("Expected not to return any deleted secrets, but got: %v", deletedSecret))
  }

  filter := ""
  listedSecrets := List(storage.Read(password(), vaultPath), filter)

  expectedListedSecrets := `[
[0da52a01-302d-4e2f-8200-a4d4226699af]
secret-1-name
secret-1-value 
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-name
secret-2-value 
[1f947c00-211c-4216-865e-1ca1fcbce693]
secret-3-name
secret-3-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %v, but got %v", expectedListedSecrets, listedSecrets))
  }
}

func TestEdit(t *testing.T) {
  vaultPath := prepareVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-c70b-4a81-87db-7839d16f0176")
  if editedSecret, err := Edit(id, "secret-2-new-name", "secret-2-new-value", vaultPath, password()); err != nil {
    t.Fatal(err)
  } else {
    expectedEditedSecret := secret.Secret{id, "secret-2-new-name", "secret-2-new-value"}
    if fmt.Sprintf("%v", expectedEditedSecret) != fmt.Sprintf("%v", editedSecret) {
      t.Fatal(fmt.Sprintf("Expected edited secret to be %s, but got %s", expectedEditedSecret, editedSecret))
    }
  }

  filter := ""
  listedSecrets := List(storage.Read(password(), vaultPath), filter)

  expectedListedSecrets := `[
[0da52a01-302d-4e2f-8200-a4d4226699af]
secret-1-name
secret-1-value 
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-new-name
secret-2-new-value 
[1f947c00-211c-4216-865e-1ca1fcbce693]
secret-3-name
secret-3-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %v, but got %v", expectedListedSecrets, listedSecrets))
  }
}

func TestEdit_NonExistingId(t *testing.T) {
  vaultPath := prepareVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-0000-0000-87db-7839d16f0176")
  expectedError := "Secret by specified id not found!"
  if editedSecret, err := Edit(id, "secret-2-new-name", "secret-2-new-value", vaultPath, password()); err.Error() != expectedError {
    t.Fatal("Expected to return an error %s, but got %s", expectedError, err.Error())
  } else if editedSecret != nil {
    t.Fatal(fmt.Sprintf("Expected not to return any edited secrets, but got: %v", editedSecret))
  }

  filter := ""
  listedSecrets := List(storage.Read(password(), vaultPath), filter)

  expectedListedSecrets := `[
[0da52a01-302d-4e2f-8200-a4d4226699af]
secret-1-name
secret-1-value 
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-name
secret-2-value 
[1f947c00-211c-4216-865e-1ca1fcbce693]
secret-3-name
secret-3-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %v, but got %v", expectedListedSecrets, listedSecrets))
  }
}

func TestChangePassword(t *testing.T) {
  vaultPath := prepareVault(t)
  defer os.Remove(vaultPath)

  newPassword := []byte("new password")
  newPasswordConfirmation := newPassword

  if err := ChangePassword(vaultPath, password(), newPassword, newPasswordConfirmation); err != nil {
    t.Fatal(err)
  }

  filter := ""
  listedSecrets := List(storage.Read(newPassword, vaultPath), filter)

  expectedListedSecrets := `[
[0da52a01-302d-4e2f-8200-a4d4226699af]
secret-1-name
secret-1-value 
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-name
secret-2-value 
[1f947c00-211c-4216-865e-1ca1fcbce693]
secret-3-name
secret-3-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %v, but got %v", expectedListedSecrets, listedSecrets))
  }
}

func TestChangePassword_ConfirmationPasswordDoesNotMatch(t *testing.T) {
  vaultPath := prepareVault(t)
  defer os.Remove(vaultPath)

  newPassword := []byte("new password")
  newPasswordConfirmation := []byte("different new password")

  expectedError := "Passwords do not match!"
  if err := ChangePassword(vaultPath, password(), newPassword, newPasswordConfirmation); err.Error() != expectedError {
    t.Fatal("Expected to return an error %s, but got %s", expectedError, err.Error())
  }

  filter := ""
  listedSecrets := List(storage.Read(password(), vaultPath), filter)

  expectedListedSecrets := `[
[0da52a01-302d-4e2f-8200-a4d4226699af]
secret-1-name
secret-1-value 
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-name
secret-2-value 
[1f947c00-211c-4216-865e-1ca1fcbce693]
secret-3-name
secret-3-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %v, but got %v", expectedListedSecrets, listedSecrets))
  }
}

func prepareVault(t *testing.T) string {
  vaultPath, err := ioutil.TempFile("", "test-vault")
  if err != nil {
    t.Fatal(err)
  }
  vaultPathStr := vaultPath.Name()

  if testVaultData, err := ioutil.ReadFile("vault_test_input.json"); err != nil {
    t.Fatal(err)
  } else {
    if err := ioutil.WriteFile(vaultPathStr, testVaultData, 0600); err != nil {
      t.Fatal(err)
    }
  }

  return vaultPathStr
}

func password() []byte {
  return []byte("secret password")
}
