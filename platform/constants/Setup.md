# Azure Subscriptions Integration Setup Guide for opencomply

This guide provides step-by-step instructions to integrate your Azure subscriptions with opencomply by creating a Service Principal with read-only access. This integration enables opencomply to provide visibility and governance capabilities over your Azure resources.

## Table of Contents

- [Azure Subscriptions Integration Setup Guide for opencomply](#azure-subscriptions-integration-setup-guide-for-opencomply)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Steps](#steps)
    - [1. Clone the Integration Scripts Repository](#1-clone-the-integration-scripts-repository)
    - [2. Run the Reader Role Assignment Script](#2-run-the-reader-role-assignment-script)
    - [3. Setup opencomply](#3-setup-opencomply)
      - [Navigate to the opencomply Dashboard:](#navigate-to-the-opencomply-dashboard)
      - [Access the Integrations Section:](#access-the-integrations-section)
      - [Enter the Required Details:](#enter-the-required-details)
      - [Complete the Integration:](#complete-the-integration)

## Prerequisites

Before you begin, ensure the following prerequisites are met:

- **Azure CLI Installed and Authenticated**: The Azure CLI must be installed on your machine and authenticated with sufficient privileges.
  - **Install Azure CLI**: Follow the [Azure CLI Installation Guide](https://learn.microsoft.com/cli/azure/install-azure-cli) to install the Azure CLI.
  - **Authenticate**: Run the following command and follow the prompts to authenticate:

    ```bash
    az login
    ```

- **opencomply Installed and Running**: Ensure that opencomply is installed and operational. Refer to the [opencomply Installation Documentation](https://github.com/opengovern/integration-automation-scripts) if needed.

## Steps

Follow the steps below to set up the Azure subscriptions integration with opencomply.

### 1. Clone the Integration Scripts Repository

The integration scripts automate the creation of the Service Principal (SPN) and role assignment.

```bash
# Clone the repository
git clone https://github.com/opengovern/integration-automation-scripts.git
# Navigate to the Azure directory
cd integration-automation-scripts/azure-subscriptions
```

### 2. Run the Reader Role Assignment Script

Execute the script to create a Service Principal (SPN) and assign it the Reader role across all your Azure subscriptions.

- Make the Script Executable (if not already):

  ```bash
  chmod +x assign_reader_role.sh
  ```

- Run the Script:

  ```bash
  ./assign_reader_role.sh
  ```

The script will perform the following actions:

- Create a Service Principal with the necessary permissions.
- Assign the Reader role to the Service Principal for each Azure subscription in your tenant.

**Note**: Ensure you have the necessary permissions to create Service Principals and assign roles in your Azure tenant.

### 3. Setup opencomply

After running the script, it will output essential details required for configuring opencomply:

- Tenant ID
- Application (Client) ID
- Object ID
- Client Secret

Use the credentials obtained to configure Azure integration within opencomply.

#### Navigate to the opencomply Dashboard:

Open your web browser and go to the opencomply portal. Log in with your administrator credentials.

#### Access the Integrations Section:

- In the sidebar, click on Integrations.
- Select Azure from the list of available integrations.
- Click on Add New Integration and choose New SPN from the options.

#### Enter the Required Details:

In the integration wizard, provide the following details:

- **Tenant ID**: Enter the Tenant ID obtained from the script output.
- **Application (Client) ID**: Enter the Application (Client) ID.
- **Client Secret**: Enter the Client Secret. Ensure this is stored securely.
- **Object ID**: Enter the Object ID associated with the Service Principal.

#### Complete the Integration:

Follow the on-screen instructions to complete the integration process. Once completed, your Azure subscriptions will be linked with opencomply, providing enhanced visibility and governance over your Azure resources.