# Microsoft Entra ID Integration Setup Guide for opencomply

This guide provides step-by-step instructions to integrate Azure Entra ID with opencomply by creating a Service Principal (SPN) with the necessary API permissions. This integration enables opencomply to provide comprehensive governance and compliance capabilities over your Azure Entra ID resources.

## Table of Contents

- [Microsoft Entra ID Integration Setup Guide for opencomply](#microsoft-entra-id-integration-setup-guide-for-opencomply)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Steps](#steps)
    - [1. Create a Service Principal in Azure Entra ID](#1-create-a-service-principal-in-azure-entra-id)
      - [Sign in to the Azure Portal](#sign-in-to-the-azure-portal)
      - [Navigate to Azure Active Directory](#navigate-to-azure-active-directory)
      - [Register a New Application](#register-a-new-application)
    - [2. Assign API Permissions to the Service Principal](#2-assign-api-permissions-to-the-service-principal)
      - [Navigate to API Permissions](#navigate-to-api-permissions)
      - [Add Permissions](#add-permissions)
      - [Select Required Permissions](#select-required-permissions)
    - [3. Grant Admin Consent for the Permissions](#3-grant-admin-consent-for-the-permissions)
      - [Grant Consent](#grant-consent)
    - [4. Obtain Required Credentials](#4-obtain-required-credentials)
      - [Generate a Client Secret](#generate-a-client-secret)
      - [Collect Application Details](#collect-application-details)
    - [5. Configure opencomply](#5-configure-opencomply)

## Prerequisites

Before you begin, ensure the following prerequisites are met:

- **Azure Portal Access**: You must have access to the Azure portal with sufficient privileges to create applications and assign permissions in Azure Entra ID (formerly Azure Active Directory).
- **Global Administrator Role**: You may need Global Administrator privileges to grant admin consent for API permissions.
- **opencomply Access**: Ensure that you have administrator access to the opencomply portal.

## Steps

Follow the steps below to set up the Azure Entra ID integration with opencomply.

### 1. Create a Service Principal in Azure Entra ID

#### Sign in to the Azure Portal

- Go to [https://portal.azure.com](https://portal.azure.com) and sign in with your administrator account.

#### Navigate to Azure Active Directory

- In the left-hand navigation pane, select **Azure Active Directory**.

#### Register a New Application

- Under **Manage**, select **App registrations**.
- Click on **New registration**.
  - **Name**: Enter a name for the application (e.g., opencomply Integration).
  - **Supported account types**: Select **Accounts in this organizational directory only**.
  - **Redirect URI**: Leave this field blank.
- Click **Register**.

### 2. Assign API Permissions to the Service Principal

#### Navigate to API Permissions

- In your application's overview page, select **API permissions** from the left-hand menu.

#### Add Permissions

- Click on **Add a permission**.
- Choose **Microsoft Graph**.
- Select **Application permissions**.

#### Select Required Permissions

Search for and select the following permissions:

- Application.Read.All
- AuditLog.Read.All
- Directory.Read.All
- Domain.Read.All
- Group.Read.All
- IdentityProvider.Read.All
- Policy.Read.All
- User.Read.All

After selecting all the permissions, click **Add permissions**.

### 3. Grant Admin Consent for the Permissions

#### Grant Consent

- Back on the API permissions page, click on **Grant admin consent for [Your Organization]**.
- Confirm by clicking **Yes**.
- Ensure that the status for each permission shows **Granted for [Your Organization]**.

### 4. Obtain Required Credentials

#### Generate a Client Secret

- Navigate to **Certificates & secrets** in the left-hand menu.
- Under **Client secrets**, click **New client secret**.
  - **Description**: Enter a description (e.g., opencomply Secret).
  - **Expires**: Select an appropriate expiration period.
- Click **Add**.

**Important**: Copy the **Value** of the client secret now; it will not be shown again.

#### Collect Application Details

- Go back to the **Overview** page of your application.
- Note down the following:
  - **Application (client) ID**
  - **Directory (tenant) ID**
  - **Object ID** (found under **Managed application in local directory**)

### 5. Configure opencomply

Use the credentials obtained to configure Azure Entra ID integration within opencomply.

- Click **Confirm** to finalize the integration.
- Wait for opencomply to validate the credentials and establish the integration.