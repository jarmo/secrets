package vault

import (
  "testing"
  "fmt"
  "os"
  "io/ioutil"
  "github.com/satori/go.uuid"
  "github.com/jarmo/secrets/secret"
  "github.com/jarmo/secrets/storage"
)

func TestListUsingArgon2idKey(t *testing.T) {
  vaultPath := prepareArgon2idVault(t)
  defer os.Remove(vaultPath)

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestListUsingScryptKey(t *testing.T) {
  vaultPath := prepareScryptVault(t)
  defer os.Remove(vaultPath)

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestList_WithFilterUsingArgon2idKey(t *testing.T) {
  vaultPath := prepareArgon2idVault(t)
  defer os.Remove(vaultPath)

  filter := "secret-2"
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

  expectedListedSecrets := `[
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-name
secret-2-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %s, but got %s", expectedListedSecrets, listedSecrets))
  }
}

func TestList_WithFilterUsingScryptKey(t *testing.T) {
  vaultPath := prepareScryptVault(t)
  defer os.Remove(vaultPath)

  filter := "secret-2"
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

  expectedListedSecrets := `[
[2b57a54a-c70b-4a81-87db-7839d16f0176]
secret-2-name
secret-2-value]`

  if fmt.Sprintf("%v", listedSecrets) != fmt.Sprintf("%v", expectedListedSecrets) {
    t.Fatal(fmt.Sprintf("Expected listed secrets to be %s, but got %s", expectedListedSecrets, listedSecrets))
  }
}

func TestAddUsingArgon2idKey(t *testing.T) {
  vaultPath := prepareArgon2idVault(t)
  defer os.Remove(vaultPath)

  addedSecret, newSecrets := Add(secrets(t, vaultPath, password()), "secret-4-name", "secret-4-value")
  storage.Write(vaultPath, password(), newSecrets)

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestAddUsingScryptKey(t *testing.T) {
  vaultPath := prepareScryptVault(t)
  defer os.Remove(vaultPath)

  addedSecret, newSecrets := Add(secrets(t, vaultPath, password()), "secret-4-name", "secret-4-value")
  storage.Write(vaultPath, password(), newSecrets)

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestDeleteUsingArgon2idKey(t *testing.T) {
  vaultPath := prepareArgon2idVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-c70b-4a81-87db-7839d16f0176")
  if deletedSecret, newSecrets, err := Delete(secrets(t, vaultPath, password()), id); err != nil {
    t.Fatal(err)
  } else {
    expectedDeletedSecret := secret.Secret{id, "secret-2-name", "secret-2-value"}
    if fmt.Sprintf("%v", expectedDeletedSecret) != fmt.Sprintf("%v", deletedSecret) {
      t.Fatal(fmt.Sprintf("Expected deleted secret to be %s, but got %s", expectedDeletedSecret, deletedSecret))
    }
    storage.Write(vaultPath, password(), newSecrets)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestDeleteUsingScryptKey(t *testing.T) {
  vaultPath := prepareScryptVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-c70b-4a81-87db-7839d16f0176")
  if deletedSecret, newSecrets, err := Delete(secrets(t, vaultPath, password()), id); err != nil {
    t.Fatal(err)
  } else {
    expectedDeletedSecret := secret.Secret{id, "secret-2-name", "secret-2-value"}
    if fmt.Sprintf("%v", expectedDeletedSecret) != fmt.Sprintf("%v", deletedSecret) {
      t.Fatal(fmt.Sprintf("Expected deleted secret to be %s, but got %s", expectedDeletedSecret, deletedSecret))
    }
    storage.Write(vaultPath, password(), newSecrets)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestDelete_NonExistingIdUsingArgon2idKey(t *testing.T) {
  vaultPath := prepareArgon2idVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-0000-0000-87db-7839d16f0176")
  expectedError := "Secret by specified id not found!"
  if deletedSecret, newSecrets, err := Delete(secrets(t, vaultPath, password()), id); err.Error() != expectedError {
    t.Fatalf("Expected to return an error %s, but got %s", expectedError, err.Error())
  } else if deletedSecret != nil {
    t.Fatal(fmt.Sprintf("Expected not to return any deleted secrets, but got: %v", deletedSecret))
  } else {
    storage.Write(vaultPath, password(), newSecrets)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestDelete_NonExistingIdUsingScryptKey(t *testing.T) {
  vaultPath := prepareScryptVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-0000-0000-87db-7839d16f0176")
  expectedError := "Secret by specified id not found!"
  if deletedSecret, newSecrets, err := Delete(secrets(t, vaultPath, password()), id); err.Error() != expectedError {
    t.Fatalf("Expected to return an error %s, but got %s", expectedError, err.Error())
  } else if deletedSecret != nil {
    t.Fatal(fmt.Sprintf("Expected not to return any deleted secrets, but got: %v", deletedSecret))
  } else {
    storage.Write(vaultPath, password(), newSecrets)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestEditUsingArgon2idKey(t *testing.T) {
  vaultPath := prepareArgon2idVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-c70b-4a81-87db-7839d16f0176")
  if editedSecret, newSecrets, err := Edit(secrets(t, vaultPath, password()), id, "secret-2-new-name", "secret-2-new-value"); err != nil {
    t.Fatal(err)
  } else {
    expectedEditedSecret := secret.Secret{id, "secret-2-new-name", "secret-2-new-value"}
    if fmt.Sprintf("%v", expectedEditedSecret) != fmt.Sprintf("%v", editedSecret) {
      t.Fatal(fmt.Sprintf("Expected edited secret to be %s, but got %s", expectedEditedSecret, editedSecret))
    }

    storage.Write(vaultPath, password(), newSecrets)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestEditUsingScryptKey(t *testing.T) {
  vaultPath := prepareScryptVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-c70b-4a81-87db-7839d16f0176")
  if editedSecret, newSecrets, err := Edit(secrets(t, vaultPath, password()), id, "secret-2-new-name", "secret-2-new-value"); err != nil {
    t.Fatal(err)
  } else {
    expectedEditedSecret := secret.Secret{id, "secret-2-new-name", "secret-2-new-value"}
    if fmt.Sprintf("%v", expectedEditedSecret) != fmt.Sprintf("%v", editedSecret) {
      t.Fatal(fmt.Sprintf("Expected edited secret to be %s, but got %s", expectedEditedSecret, editedSecret))
    }

    storage.Write(vaultPath, password(), newSecrets)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestEdit_NonExistingIdUsingArgon2idKey(t *testing.T) {
  vaultPath := prepareArgon2idVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-0000-0000-87db-7839d16f0176")
  expectedError := "Secret by specified id not found!"
  if editedSecret, newSecrets, err := Edit(secrets(t, vaultPath, password()), id, "secret-2-new-name", "secret-2-new-value"); err.Error() != expectedError {
    t.Fatalf("Expected to return an error %s, but got %s", expectedError, err.Error())
  } else if editedSecret != nil {
    t.Fatal(fmt.Sprintf("Expected not to return any edited secrets, but got: %v", editedSecret))
  } else {
    storage.Write(vaultPath, password(), newSecrets)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestEdit_NonExistingIdUsingScryptKey(t *testing.T) {
  vaultPath := prepareScryptVault(t)
  defer os.Remove(vaultPath)

  id, _ := uuid.FromString("2b57a54a-0000-0000-87db-7839d16f0176")
  expectedError := "Secret by specified id not found!"
  if editedSecret, newSecrets, err := Edit(secrets(t, vaultPath, password()), id, "secret-2-new-name", "secret-2-new-value"); err.Error() != expectedError {
    t.Fatalf("Expected to return an error %s, but got %s", expectedError, err.Error())
  } else if editedSecret != nil {
    t.Fatal(fmt.Sprintf("Expected not to return any edited secrets, but got: %v", editedSecret))
  } else {
    storage.Write(vaultPath, password(), newSecrets)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestChangePasswordUsingArgon2idKey(t *testing.T) {
  vaultPath := prepareArgon2idVault(t)
  defer os.Remove(vaultPath)

  newPassword := []byte("new password")
  newPasswordConfirmation := newPassword

  if err := ChangePassword(vaultPath, password(), newPassword, newPasswordConfirmation); err != nil {
    t.Fatal(err)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, newPassword), filter)

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

func TestChangePasswordUsingScryptKey(t *testing.T) {
  vaultPath := prepareScryptVault(t)
  defer os.Remove(vaultPath)

  newPassword := []byte("new password")
  newPasswordConfirmation := newPassword

  if err := ChangePassword(vaultPath, password(), newPassword, newPasswordConfirmation); err != nil {
    t.Fatal(err)
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, newPassword), filter)

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

func TestChangePassword_ConfirmationPasswordDoesNotMatchUsingArgon2idKey(t *testing.T) {
  vaultPath := prepareArgon2idVault(t)
  defer os.Remove(vaultPath)

  newPassword := []byte("new password")
  newPasswordConfirmation := []byte("different new password")

  expectedError := "Passwords do not match!"
  if err := ChangePassword(vaultPath, password(), newPassword, newPasswordConfirmation); err.Error() != expectedError {
    t.Fatalf("Expected to return an error %s, but got %s", expectedError, err.Error())
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func TestChangePassword_ConfirmationPasswordDoesNotMatchUsingScryptKey(t *testing.T) {
  vaultPath := prepareScryptVault(t)
  defer os.Remove(vaultPath)

  newPassword := []byte("new password")
  newPasswordConfirmation := []byte("different new password")

  expectedError := "Passwords do not match!"
  if err := ChangePassword(vaultPath, password(), newPassword, newPasswordConfirmation); err.Error() != expectedError {
    t.Fatalf("Expected to return an error %s, but got %s", expectedError, err.Error())
  }

  filter := ""
  listedSecrets := List(secrets(t, vaultPath, password()), filter)

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

func prepareArgon2idVault(t *testing.T) string {
  return prepareVault(t, "vault_test_argon2id_input.json")
}

func prepareScryptVault(t *testing.T) string {
  return prepareVault(t, "vault_test_scrypt_input.json")
}

func prepareVault(t *testing.T, vaultDataPath string) string {
  vaultPath, err := ioutil.TempFile("", "test-vault")
  if err != nil {
    t.Fatal(err)
  }
  vaultPathStr := vaultPath.Name()

  if testVaultData, err := ioutil.ReadFile(vaultDataPath); err != nil {
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

func secrets(t *testing.T, vaultPath string, password []byte) []secret.Secret {
  secrets, err := storage.Read(vaultPath, password)
  if err != nil {
    t.Fatal(err)
  }

  return secrets
}
