// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://github.com/Azure/magic-module-specs
//
// ----------------------------------------------------------------------------

package azurerm



func resourceArmBackend() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmBackendCreate,
        Read: resourceArmBackendRead,
        Update: resourceArmBackendUpdate,
        Delete: resourceArmBackendDelete,

        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "resource_group": azure.SchemaResourceGroupNameDiffSuppress(),

            "backendid": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "protocol": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    string(apimanagement.http),
                    string(apimanagement.soap),
                }, false),
            },

            "url": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "certificate": {
                Type: schema.TypeList,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "description": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "header": {
                Type: schema.TypeMap,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "password": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "query": {
                Type: schema.TypeMap,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Schema{Type: schema.TypeString},
            },

            "resource_id": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "skip_certificate_chain_validation": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "skip_certificate_name_validation": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "title": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "username": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },
        },
    }
}

func resourceArmBackendCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backendsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    backendid := d.Get("backendid").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, resourceGroup, name, backendid)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_backend", *existing.ID)
        }
    }

    certificate := d.Get("certificate").([]interface{})
    description := d.Get("description").(string)
    header := d.Get("header").(map[string]interface{})
    password := d.Get("password").(string)
    protocol := d.Get("protocol").(string)
    query := d.Get("query").(map[string]interface{})
    resourceId := d.Get("resource_id").(string)
    skipCertificateChainValidation := d.Get("skip_certificate_chain_validation").(bool)
    skipCertificateNameValidation := d.Get("skip_certificate_name_validation").(bool)
    title := d.Get("title").(string)
    url := d.Get("url").(string)
    username := d.Get("username").(string)

    parameters := apimanagement.BackendContract{
        Certificate: utils.ExpandStringSlice(certificate),
        Description: utils.String(description),
        Header: utils.ExpandKeyValuePairs(header),
        Password: utils.String(password),
        BackendProperties: &apimanagement.BackendProperties{
            SkipCertificateChainValidation: utils.Bool(skipCertificateChainValidation),
            SkipCertificateNameValidation: utils.Bool(skipCertificateNameValidation),
        },
        Protocol: apimanagement.BackendProtocol(protocol),
        Query: utils.ExpandKeyValuePairs(query),
        ResourceID: utils.String(resourceId),
        Title: utils.String(title),
        URL: utils.String(url),
        Username: utils.String(username),
    }


    if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, backendid, parameters); err != nil {
        return fmt.Errorf("Error creating Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }


    resp, err := client.Get(ctx, resourceGroup, name, backendid)
    if err != nil {
        return fmt.Errorf("Error retrieving Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Backend %q (Backendid %q / Resource Group %q) ID", name, backendid, resourceGroup)
    }
    d.SetId(*resp.ID)

    return resourceArmBackendRead(d, meta)
}

func resourceArmBackendRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backendsClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    backendid := id.Path["backends"]

    resp, err := client.Get(ctx, resourceGroup, name, backendid)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Backend %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }


    d.Set("name", name)
    d.Set("resource_group", resourceGroup)
    d.Set("backendid", backendid)
    d.Set("certificate", utils.FlattenStringSlice(resp.Certificate))
    d.Set("description", resp.Description)
    d.Set("header", utils.FlattenKeyValuePairs(resp.Header))
    d.Set("password", resp.Password)
    d.Set("protocol", string(resp.Protocol))
    d.Set("query", utils.FlattenKeyValuePairs(resp.Query))
    d.Set("resource_id", resp.ResourceID)
    if backendProperties := resp.BackendProperties; backendProperties != nil {
        d.Set("skip_certificate_chain_validation", backendProperties.SkipCertificateChainValidation)
        d.Set("skip_certificate_name_validation", backendProperties.SkipCertificateNameValidation)
    }
    d.Set("title", resp.Title)
    d.Set("url", resp.URL)
    d.Set("username", resp.Username)

    return nil
}

func resourceArmBackendUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backendsClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group").(string)
    backendid := d.Get("backendid").(string)
    certificate := d.Get("certificate").([]interface{})
    description := d.Get("description").(string)
    header := d.Get("header").(map[string]interface{})
    password := d.Get("password").(string)
    protocol := d.Get("protocol").(string)
    query := d.Get("query").(map[string]interface{})
    resourceId := d.Get("resource_id").(string)
    skipCertificateChainValidation := d.Get("skip_certificate_chain_validation").(bool)
    skipCertificateNameValidation := d.Get("skip_certificate_name_validation").(bool)
    title := d.Get("title").(string)
    url := d.Get("url").(string)
    username := d.Get("username").(string)

    parameters := apimanagement.BackendContract{
        Certificate: utils.ExpandStringSlice(certificate),
        Description: utils.String(description),
        Header: utils.ExpandKeyValuePairs(header),
        Password: utils.String(password),
        BackendProperties: &apimanagement.BackendProperties{
            SkipCertificateChainValidation: utils.Bool(skipCertificateChainValidation),
            SkipCertificateNameValidation: utils.Bool(skipCertificateNameValidation),
        },
        Protocol: apimanagement.BackendProtocol(protocol),
        Query: utils.ExpandKeyValuePairs(query),
        ResourceID: utils.String(resourceId),
        Title: utils.String(title),
        URL: utils.String(url),
        Username: utils.String(username),
    }


    if _, err := client.Update(ctx, resourceGroup, name, backendid, parameters); err != nil {
        return fmt.Errorf("Error updating Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }

    return resourceArmBackendRead(d, meta)
}

func resourceArmBackendDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).backendsClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    resourceGroup := id.ResourceGroup
    name := id.Path["service"]
    backendid := id.Path["backends"]

    if _, err := client.Delete(ctx, resourceGroup, name, backendid); err != nil {
        return fmt.Errorf("Error deleting Backend %q (Backendid %q / Resource Group %q): %+v", name, backendid, resourceGroup, err)
    }

    return nil
}
