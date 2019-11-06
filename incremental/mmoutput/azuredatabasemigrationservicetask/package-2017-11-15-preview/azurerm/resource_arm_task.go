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



func resourceArmTask() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmTaskCreate,
        Read: resourceArmTaskRead,
        Update: resourceArmTaskUpdate,
        Delete: resourceArmTaskDelete,

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

            "name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "group_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "project_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "service_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.NoEmptyStrings,
            },

            "etag": {
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
            },

            "errors": {
                Type: schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "code": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                        "details": {
                            Type: schema.TypeList,
                            Optional: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "code": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                    "details": {
                                        Type: schema.TypeList,
                                        Optional: true,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "code": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                                "message": {
                                                    Type: schema.TypeString,
                                                    Optional: true,
                                                },
                                            },
                                        },
                                    },
                                    "message": {
                                        Type: schema.TypeString,
                                        Optional: true,
                                    },
                                },
                            },
                        },
                        "message": {
                            Type: schema.TypeString,
                            Optional: true,
                        },
                    },
                },
            },

            "state": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourceArmTaskCreate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).tasksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    groupName := d.Get("group_name").(string)
    projectName := d.Get("project_name").(string)
    serviceName := d.Get("service_name").(string)

    if features.ShouldResourcesBeImported() && d.IsNewResource() {
        existing, err := client.Get(ctx, groupName, serviceName, projectName, name)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("Error checking for present of existing Task %q (Project Name %q / Service Name %q / Group Name %q): %+v", name, projectName, serviceName, groupName, err)
            }
        }
        if existing.ID != nil && *existing.ID != "" {
            return tf.ImportAsExistsError("azurerm_task", *existing.ID)
        }
    }

    etag := d.Get("etag").(string)

    parameters := azuredatabasemigrationservice.ProjectTask{
        Etag: utils.String(etag),
    }


    if _, err := client.CreateOrUpdate(ctx, groupName, serviceName, projectName, name, parameters); err != nil {
        return fmt.Errorf("Error creating Task %q (Project Name %q / Service Name %q / Group Name %q): %+v", name, projectName, serviceName, groupName, err)
    }


    resp, err := client.Get(ctx, groupName, serviceName, projectName, name)
    if err != nil {
        return fmt.Errorf("Error retrieving Task %q (Project Name %q / Service Name %q / Group Name %q): %+v", name, projectName, serviceName, groupName, err)
    }
    if resp.ID == nil {
        return fmt.Errorf("Cannot read Task %q (Project Name %q / Service Name %q / Group Name %q) ID", name, projectName, serviceName, groupName)
    }
    d.SetId(*resp.ID)

    return resourceArmTaskRead(d, meta)
}

func resourceArmTaskRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).tasksClient
    ctx := meta.(*ArmClient).StopContext

    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    groupName := id.ResourceGroup
    serviceName := id.Path["services"]
    projectName := id.Path["projects"]
    name := id.Path["tasks"]

    resp, err := client.Get(ctx, groupName, serviceName, projectName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] Task %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("Error reading Task %q (Project Name %q / Service Name %q / Group Name %q): %+v", name, projectName, serviceName, groupName, err)
    }


    d.Set("name", name)
    d.Set("name", resp.Name)
    if projectTaskProperties := resp.ProjectTaskProperties; projectTaskProperties != nil {
        if err := d.Set("errors", flattenArmTaskODataError(projectTaskProperties.Errors)); err != nil {
            return fmt.Errorf("Error setting `errors`: %+v", err)
        }
        d.Set("state", string(projectTaskProperties.State))
    }
    d.Set("etag", resp.Etag)
    d.Set("group_name", groupName)
    d.Set("project_name", projectName)
    d.Set("service_name", serviceName)
    d.Set("type", resp.Type)

    return nil
}

func resourceArmTaskUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).tasksClient
    ctx := meta.(*ArmClient).StopContext

    name := d.Get("name").(string)
    etag := d.Get("etag").(string)
    groupName := d.Get("group_name").(string)
    projectName := d.Get("project_name").(string)
    serviceName := d.Get("service_name").(string)

    parameters := azuredatabasemigrationservice.ProjectTask{
        Etag: utils.String(etag),
    }


    if _, err := client.Update(ctx, groupName, serviceName, projectName, name, parameters); err != nil {
        return fmt.Errorf("Error updating Task %q (Project Name %q / Service Name %q / Group Name %q): %+v", name, projectName, serviceName, groupName, err)
    }

    return resourceArmTaskRead(d, meta)
}

func resourceArmTaskDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*ArmClient).tasksClient
    ctx := meta.(*ArmClient).StopContext


    id, err := azure.ParseAzureResourceID(d.Id())
    if err != nil {
        return err
    }
    groupName := id.ResourceGroup
    serviceName := id.Path["services"]
    projectName := id.Path["projects"]
    name := id.Path["tasks"]

    if _, err := client.Delete(ctx, groupName, serviceName, projectName, name); err != nil {
        return fmt.Errorf("Error deleting Task %q (Project Name %q / Service Name %q / Group Name %q): %+v", name, projectName, serviceName, groupName, err)
    }

    return nil
}


func flattenArmTaskODataError(input *[]azuredatabasemigrationservice.ODataError) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        v := make(map[string]interface{})


        results = append(results, v)
    }

    return results
}
