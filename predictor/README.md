### stock prediction model

Development guide:

Since using both Conda and Docker is complicated and leads to heavy images, for local development (especially for training models, etc) use a conda environment, and for the Docker predictor part, we are going to use a pip environment. This means that after installing more dependencies through conda, we need to convert the conda environment into a pip requirements file so we can create images off of that. Instructions are as follows:
```
# for creating the conda env
conda env create -n tradebot -f dependencies/environment.yml
conda activate tradebot

# save conda env by
conda env export > dependencies/environment.yml

# save conda env to a pip requirements.txt
pip list --format=freeze > dependencies/requirements.txt
```
Considerations:
- Might need to change to a base Python docker image and then install PyTorch