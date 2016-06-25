package content

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/consul/vendor/github.com/hashicorp/go-uuid"
    "os"
)

func createContentFile(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    resourceData.SetId(id)

    if err = _saveContentFile(resourceData); err != nil {
        return err
    }

    return nil
}

func readContentFile(resourceData *schema.ResourceData, _ interface{}) error {
    if err := _saveContentFile(resourceData); err != nil {
        return err
    }

    return nil
}

func updateContentFile(resourceData *schema.ResourceData, _ interface{}) error {
    if !resourceData.HasChange("content") {
        return nil
    }

    if err := _saveContentFile(resourceData); err != nil {
        return err
    }

    return nil
}

func deleteContentFile(resourceData *schema.ResourceData, _ interface{}) error {
    if err := os.Remove(resourceData.Get("path").(string)); err != nil {
        return err
    }

    resourceData.SetId("")
    resourceData.Set("file", "")

    return nil
}

func contentFileExists(resourceData *schema.ResourceData, _ interface{}) (bool, error) {
    _, err := os.Stat(resourceData.Get("path").(string))
    if err != nil && os.IsNotExist(err) {
        return false, nil
    }

    return err == nil, err
}

func _saveContentFile(resourceData *schema.ResourceData) error {
    path := resourceData.Get("path").(string)
    content := resourceData.Get("content").(string)

    file, err := os.Create(path)
    if err != nil {
        return err
    }

    defer file.Close()

    if _, err = file.WriteString(content); err != nil {
        return err
    }

    if err = file.Sync(); err != nil {
        return err
    }

    resourceData.Set("file", path)

    return nil
}
