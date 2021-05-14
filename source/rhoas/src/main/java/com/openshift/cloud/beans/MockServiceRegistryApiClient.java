package com.openshift.cloud.beans;

import java.util.ArrayList;
import java.util.List;

import javax.enterprise.context.ApplicationScoped;
import javax.enterprise.inject.Alternative;

import com.openshift.cloud.api.srs.models.RegistryRest;
import com.openshift.cloud.controllers.ConditionAwareException;

@ApplicationScoped
@Alternative
public class MockServiceRegistryApiClient extends ServiceRegistryApiClient {

  public List<RegistryRest> listRegistries(String accessToken) throws ConditionAwareException {
    return new ArrayList<>();
  }

  public RegistryRest getServiceRegistryById(String registryId, String accessToken)
      throws ConditionAwareException {
    return new RegistryRest();
  }

}
